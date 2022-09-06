//go:generate mockgen -source=./storage.go -destination=./mocks/storage.go -package=mocks
package storage

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
)

type AppleMutableStorage interface {
	Add(ctx context.Context, entity *applePkg.Apple) error
	Update(ctx context.Context, entity *applePkg.Apple) error
	Delete(ctx context.Context, id uint64) error
}

type AppleReadableStorage interface {
	List(ctx context.Context, opts *PaginationOpts) ([]applePkg.Apple, error)
	Get(ctx context.Context, id uint64) (*applePkg.Apple, error)
}

type AppleStorage interface {
	AppleMutableStorage
	AppleReadableStorage
}
