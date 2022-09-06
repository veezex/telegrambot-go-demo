package monitoring

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type storageMonitoring struct {
	storage storage.AppleStorage
}

func NewStorageMonitoring(storage storage.AppleStorage) storage.AppleStorage {
	return &storageMonitoring{
		storage: storage,
	}
}

func (m *storageMonitoring) Add(ctx context.Context, entity *applePkg.Apple) error {
	err := m.storage.Add(ctx, entity)
	inRequests.Inc()

	if err != nil {
		failedRequests.Inc()
		return err
	}

	successRequests.Inc()
	return nil
}

func (m *storageMonitoring) Update(ctx context.Context, entity *applePkg.Apple) error {
	err := m.storage.Update(ctx, entity)
	inRequests.Inc()

	if err != nil {
		failedRequests.Inc()
		return err
	}

	successRequests.Inc()
	return nil
}

func (m *storageMonitoring) Delete(ctx context.Context, id uint64) error {
	err := m.storage.Delete(ctx, id)
	inRequests.Inc()

	if err != nil {
		failedRequests.Inc()
		return err
	}

	successRequests.Inc()
	return nil
}

func (m *storageMonitoring) List(ctx context.Context, opts *storage.PaginationOpts) ([]applePkg.Apple, error) {
	apples, err := m.storage.List(ctx, opts)
	inRequests.Inc()

	if err != nil {
		failedRequests.Inc()
		return nil, err
	}

	successRequests.Inc()
	return apples, nil
}

func (m *storageMonitoring) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	apple, err := m.storage.Get(ctx, id)
	inRequests.Inc()

	if err != nil {
		failedRequests.Inc()
		return nil, err
	}

	successRequests.Inc()
	return apple, nil
}
