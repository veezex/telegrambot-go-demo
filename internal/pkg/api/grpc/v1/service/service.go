package service

import (
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
)

type impl struct {
	pb.UnimplementedAppleServiceServer
	stor storagePkg.AppleStorage
}

func New(stor storagePkg.AppleStorage) pb.AppleServiceServer {
	return &impl{
		stor: stor,
	}
}
