package client

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
	"strings"

	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type remote struct {
	client pb.AppleServiceClient
}

func New(client pb.AppleServiceClient) storage.AppleReadableStorage {
	return &remote{
		client: client,
	}
}

func (r *remote) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	_, err := r.client.AppleGet(ctx, &pb.AppleGetRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *remote) List(ctx context.Context, opts *storagePkg.PaginationOpts) ([]applePkg.Apple, error) {
	request := &pb.AppleListRequest{}
	if opts != nil {
		if strings.ToUpper(opts.Order) == "DESC" {
			request.Order = pb.SortOrder_SORT_ORDER_DESC
		} else {
			request.Order = pb.SortOrder_SORT_ORDER_ASC
		}

		request.Offset = opts.Offset
		request.Limit = opts.Limit
	} else {
		request.Order = pb.SortOrder_SORT_ORDER_ASC
		request.Offset = 0
		request.Limit = 10
	}

	_, err := r.client.AppleList(ctx, request)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
