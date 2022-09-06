package service

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *impl) AppleGet(ctx context.Context, in *pb.AppleGetRequest) (*pb.AppleGetResponse, error) {
	apple, err := i.stor.Get(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, storage.ErrNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AppleGetResponse{
		Color:   apple.Color.Name,
		ColorId: apple.Color.Id,
		Price:   apple.Price,
		Id:      apple.Id,
	}, nil
}
