package combined

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type combined struct {
	writeClient storage.AppleMutableStorage
	readClient  storage.AppleReadableStorage
}

func New(writeClient storage.AppleMutableStorage, readClient storage.AppleReadableStorage) storage.AppleStorage {
	return &combined{
		writeClient: writeClient,
		readClient:  readClient,
	}
}

func (c *combined) Add(ctx context.Context, entity *applePkg.Apple) error {
	return c.writeClient.Add(ctx, entity)
}

func (c *combined) Update(ctx context.Context, entity *applePkg.Apple) error {
	return c.writeClient.Update(ctx, entity)
}

func (c *combined) Delete(ctx context.Context, id uint64) error {
	return c.writeClient.Delete(ctx, id)
}

func (c *combined) List(ctx context.Context, opts *storage.PaginationOpts) ([]applePkg.Apple, error) {
	return c.readClient.List(ctx, opts)
}

func (c *combined) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	return c.readClient.Get(ctx, id)
}
