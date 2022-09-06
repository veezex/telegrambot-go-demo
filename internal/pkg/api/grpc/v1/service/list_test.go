package service

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("can get a list", func(t *testing.T) {
		f := setUp(t)

		apple := applePkg.Apple{
			Color: color.Color{
				Name: fake().Color,
			},
			Price: fake().Price,
		}

		var limit uint64 = 225
		var offset uint64 = 124
		order := pb.SortOrder_SORT_ORDER_DESC

		orderStr := "asc"
		if order == pb.SortOrder_SORT_ORDER_DESC {
			orderStr = "desc"
		}

		stream := newStreamListMock()

		f.storage.EXPECT().
			List(gomock.Any(), &storage.PaginationOpts{
				Offset: offset,
				Limit:  limit,
				Order:  orderStr,
			}).
			Return([]applePkg.Apple{apple}, nil).
			Times(1)

		err := f.api.AppleList(&pb.AppleListRequest{
			Offset: offset,
			Limit:  limit,
			Order:  order,
		}, stream)

		require.NoError(t, err)
		list := stream.extractList()

		assert.Equal(t, []applePkg.Apple{apple}, list)
	})

	t.Run("throws an internal error", func(t *testing.T) {
		f := setUp(t)

		stream := newStreamListMock()

		f.storage.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("custom error")).
			Times(1)

		err := f.api.AppleList(&pb.AppleListRequest{}, stream)

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.Internal)
	})
}
