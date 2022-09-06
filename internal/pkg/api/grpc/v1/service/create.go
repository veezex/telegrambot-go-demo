package service

import (
	"context"
	"github.com/pkg/errors"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *impl) AppleCreate(ctx context.Context, in *pb.AppleCreateRequest) (*pb.AppleCreateResponse, error) {
	a := applePkg.Apple{
		Color: colorPkg.Color{
			Name: in.GetColor(),
		},
		Price: in.GetPrice(),
	}

	if err := i.stor.Add(ctx, &a); err != nil {
		if errors.Is(err, storage.ErrExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AppleCreateResponse{
		Id:      a.Id,
		ColorId: a.Color.Id,
	}, nil
}
