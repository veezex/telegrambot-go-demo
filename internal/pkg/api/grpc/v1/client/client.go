package client

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	"gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"io"
	"strings"

	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type remote struct {
	client v1.AppleServiceClient
}

func New(client v1.AppleServiceClient) storage.AppleStorage {
	return &remote{
		client: client,
	}
}

func (r *remote) Add(ctx context.Context, a *applePkg.Apple) error {
	response, err := r.client.AppleCreate(ctx, &v1.AppleCreateRequest{
		Color: a.Color.Name,
		Price: a.Price,
	})

	if err != nil {
		return err
	}

	a.Id = response.GetId()
	a.Color.Id = response.GetColorId()

	return nil
}

func (r *remote) Delete(ctx context.Context, id uint64) error {
	_, err := r.client.AppleDelete(ctx, &v1.AppleDeleteRequest{
		Id: id,
	})

	return err
}

func (r *remote) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	response, err := r.client.AppleGet(ctx, &v1.AppleGetRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return mapPbApple(response), nil
}

func (r *remote) List(ctx context.Context, opts *storagePkg.PaginationOpts) ([]applePkg.Apple, error) {
	request := &v1.AppleListRequest{}
	if opts != nil {
		if strings.ToUpper(opts.Order) == "DESC" {
			request.Order = v1.SortOrder_SORT_ORDER_DESC
		} else {
			request.Order = v1.SortOrder_SORT_ORDER_ASC
		}

		request.Offset = opts.Offset
		request.Limit = opts.Limit
	} else {
		request.Order = v1.SortOrder_SORT_ORDER_ASC
		request.Offset = 0
		request.Limit = 10
	}

	stream, err := r.client.AppleList(ctx, request)

	if err != nil {
		return nil, err
	}

	apples := make([]applePkg.Apple, 0)
	for {
		item, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return apples, nil
			}
			return nil, err
		}

		apples = append(apples, *mapPbApple(item))
	}
}

func (r *remote) Update(ctx context.Context, a *applePkg.Apple) error {
	result, err := r.client.AppleUpdate(ctx, &v1.AppleUpdateRequest{
		Id:    a.Id,
		Color: a.Color.Name,
		Price: a.Price,
	})

	if err != nil {
		return err
	}

	a.Color.Id = result.ColorId
	return nil
}

func mapPbApple(input *v1.AppleGetResponse) *applePkg.Apple {
	return &applePkg.Apple{
		Id: input.GetId(),
		Color: colorPkg.Color{
			Id:   input.GetColorId(),
			Name: input.GetColor(),
		},
		Price: input.GetPrice(),
	}
}
