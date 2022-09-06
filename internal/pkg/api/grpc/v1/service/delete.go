package service

import (
	"context"
	"errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *impl) AppleDelete(ctx context.Context, in *pb.AppleDeleteRequest) (*pb.AppleDeleteResponse, error) {
	if err := i.stor.Delete(ctx, in.GetId()); err != nil {
		if errors.Is(err, storage.ErrNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AppleDeleteResponse{}, nil
}
