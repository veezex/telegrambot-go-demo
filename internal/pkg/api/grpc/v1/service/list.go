package service

import (
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	"gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *impl) AppleList(in *v1.AppleListRequest, stream v1.AppleService_AppleListServer) error {
	order := "asc"
	if in.GetOrder() == v1.SortOrder_SORT_ORDER_DESC {
		order = "desc"
	}

	apples, err := i.stor.List(stream.Context(), &storage.PaginationOpts{
		Order:  order,
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, apple := range apples {
		err := stream.Send(&v1.AppleGetResponse{
			Color:   apple.Color.Name,
			ColorId: apple.Color.Id,
			Price:   apple.Price,
			Id:      apple.Id,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
