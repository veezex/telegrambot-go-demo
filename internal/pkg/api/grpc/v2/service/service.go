package service

import (
	"context"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
)

type ResultHandler = func(ctx context.Context, value []byte, meta interface{}) error

type impl struct {
	pb.UnimplementedAppleServiceServer
	stor          storagePkg.AppleStorage
	publishResult ResultHandler
}

func New(stor storagePkg.AppleStorage, resultFunc ResultHandler) pb.AppleServiceServer {
	return &impl{
		stor:          stor,
		publishResult: resultFunc,
	}
}
