package throttled

import (
	"context"
	"github.com/pkg/errors"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/rate_limiter"
)

var (
	ErrUnknown = errors.New("Unknown entity")
)

type throttled struct {
	limiter rate_limiter.RateLimiter
	stor    storagePkg.AppleStorage
}

func New(stor storagePkg.AppleStorage, limiter rate_limiter.RateLimiter) storagePkg.AppleStorage {
	return &throttled{
		stor:    stor,
		limiter: limiter,
	}
}

func (t *throttled) Add(ctx context.Context, a *applePkg.Apple) error {
	_, err := t.limiter.Run(func() (interface{}, error) {
		return nil, t.stor.Add(ctx, a)
	})

	return err
}

func (t *throttled) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	result, err := t.limiter.Run(func() (interface{}, error) {
		return t.stor.Get(ctx, id)
	})

	if err != nil {
		return nil, err
	}

	entity, ok := result.(*applePkg.Apple)
	if !ok {
		return nil, errors.Wrapf(ErrUnknown, "<%v>", result)
	}

	return entity, nil
}

func (t *throttled) List(ctx context.Context, opts *storagePkg.PaginationOpts) ([]applePkg.Apple, error) {
	result, err := t.limiter.Run(func() (interface{}, error) {
		return t.stor.List(ctx, opts)
	})

	if err != nil {
		return nil, err
	}

	entities, ok := result.([]applePkg.Apple)
	if !ok {
		return nil, errors.Wrapf(ErrUnknown, "<%v>", result)
	}

	return entities, nil
}

func (t *throttled) Update(ctx context.Context, a *applePkg.Apple) error {
	_, err := t.limiter.Run(func() (interface{}, error) {
		return nil, t.stor.Update(ctx, a)
	})

	return err
}

func (t *throttled) Delete(ctx context.Context, id uint64) error {
	_, err := t.limiter.Run(func() (interface{}, error) {
		return nil, t.stor.Delete(ctx, id)
	})

	return err
}
