package service

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
)

func (i *impl) AppleList(_ context.Context, in *pb.AppleListRequest) (*pb.AppleListResponse, error) {
	inOrder := in.GetOrder()
	inLimit := in.GetLimit()
	inOffset := in.GetOffset()

	go func() {
		order := "asc"
		if inOrder == pb.SortOrder_SORT_ORDER_DESC {
			order = "desc"
		}

		bgCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		apples, err := i.stor.List(bgCtx, &storage.PaginationOpts{
			Order:  order,
			Limit:  inLimit,
			Offset: inOffset,
		})
		if err != nil {
			logrus.Error(err)
			return
		}

		applesStr, err := json.Marshal(apples)
		if err != nil {
			logrus.Error(err)
			return
		}

		err = i.publishResult(bgCtx, applesStr, "list")
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	return &pb.AppleListResponse{}, nil
}
