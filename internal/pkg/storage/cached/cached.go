package cached

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/cache"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

var (
	errMalformatted = errors.New("Malformatted error")
)

type cached struct {
	cache    cache.Cacher
	storage  storage.AppleStorage
	listKeys map[string]struct{}
}

func New(storage storage.AppleStorage, cache cache.Cacher) storage.AppleStorage {
	return &cached{
		storage:  storage,
		cache:    cache,
		listKeys: make(map[string]struct{}),
	}
}

func (c *cached) Add(ctx context.Context, a *applePkg.Apple) error {
	if err := c.storage.Add(ctx, a); err != nil {
		return err
	}

	if err := c.invalidate(a.Id); err != nil {
		return err
	}

	return nil
}

func (c *cached) Delete(ctx context.Context, id uint64) error {
	if err := c.storage.Delete(ctx, id); err != nil {
		return err
	}

	if err := c.invalidate(id); err != nil {
		return err
	}

	return nil
}

func (c *cached) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	genRes, err := c.cache.Get(makeKey(id), func() (cache.CacheValue, error) {
		apple, err := c.storage.Get(ctx, id)
		if err != nil {
			return nil, err
		}

		if apple == nil {
			return nil, nil
		}

		appleStr, err := json.Marshal(*apple)
		if err != nil {
			return nil, err
		}

		return appleStr, nil
	})

	if err != nil {
		return nil, err
	}

	if genRes == nil {
		return nil, nil

	}

	var val applePkg.Apple
	err = json.Unmarshal(genRes, &val)
	if err != nil {
		return nil, err
	}

	return &val, nil
}

func (c *cached) List(ctx context.Context, opts *storagePkg.PaginationOpts) ([]applePkg.Apple, error) {
	// let's cache only first page and skip others
	if opts != nil && opts.Offset != 0 {
		return c.storage.List(ctx, opts)
	}

	key := "list_full"
	if opts != nil {
		key = fmt.Sprintf("list_%d_%d_%s", opts.Limit, opts.Offset, opts.Order)
	}

	c.listKeys[key] = struct{}{}

	genRes, err := c.cache.Get(key, func() (cache.CacheValue, error) {
		list, err := c.storage.List(ctx, opts)

		listStr, err := json.Marshal(list)
		if err != nil {
			return nil, err
		}

		return listStr, nil
	})

	if err != nil {
		return nil, err
	}

	var val []applePkg.Apple
	err = json.Unmarshal(genRes, &val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *cached) Update(ctx context.Context, a *applePkg.Apple) error {
	if err := c.storage.Update(ctx, a); err != nil {
		return err
	}

	if err := c.invalidate(a.Id); err != nil {
		return err
	}

	return nil
}

func (c *cached) invalidate(id uint64) error {
	if err := c.cache.Invalidate(makeKey(id)); err != nil {
		return err
	}

	for key, _ := range c.listKeys {
		if err := c.cache.Invalidate(key); err != nil {
			return err
		}
	}

	c.listKeys = make(map[string]struct{})
	return nil
}

func makeKey(id uint64) string {
	return fmt.Sprintf("key_%d", id)
}
