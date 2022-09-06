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

func (i *impl) AppleUpdate(ctx context.Context, in *pb.AppleUpdateRequest) (*pb.AppleUpdateResponse, error) {
	a := applePkg.Apple{
		Color: colorPkg.Color{
			Name: in.GetColor(),
		},
		Price: in.GetPrice(),
		Id:    in.GetId(),
	}

	if err := i.stor.Update(ctx, &a); err != nil {
		if errors.Is(err, storage.ErrNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AppleUpdateResponse{
		ColorId: a.Color.Id,
	}, nil
}
