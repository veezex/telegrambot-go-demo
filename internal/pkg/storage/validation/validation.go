package validation

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type validation struct {
	stor storage.AppleStorage
}

func New(stor storage.AppleStorage) storage.AppleStorage {
	return &validation{
		stor: stor,
	}
}

func (v *validation) Add(ctx context.Context, a *applePkg.Apple) error {
	if err := applePkg.Validate(a); err != nil {
		return err
	}

	return v.stor.Add(ctx, a)
}

func (v *validation) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	if id < 1 {
		return nil, errors.Wrapf(entities.ErrValidation, "id can't be less than 1, <%d>", id)
	}

	return v.stor.Get(ctx, id)
}

func (v *validation) List(ctx context.Context, opts *storagePkg.PaginationOpts) ([]applePkg.Apple, error) {
	if opts != nil {
		if err := opts.Validate(); err != nil {
			return nil, err
		}
	}

	return v.stor.List(ctx, opts)
}

func (v *validation) Update(ctx context.Context, a *applePkg.Apple) error {
	if err := applePkg.Validate(a); err != nil {
		return err
	}

	return v.stor.Update(ctx, a)
}

func (v *validation) Delete(ctx context.Context, id uint64) error {
	if id < 1 {
		return errors.Wrapf(entities.ErrValidation, "id can't be less than 1, <%d>", id)
	}

	return v.stor.Delete(ctx, id)
}
