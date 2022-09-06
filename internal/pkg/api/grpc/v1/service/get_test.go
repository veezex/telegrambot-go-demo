package service

import (
	"context"
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

func TestGet(t *testing.T) {
	t.Run("can get an apple", func(t *testing.T) {
		f := setUp(t)

		apple := applePkg.Apple{
			Color: color.Color{
				Name: fake().Color,
			},
			Price: fake().Price,
		}

		f.storage.EXPECT().
			Get(gomock.Any(), apple.Id).
			Return(&apple, nil).
			Times(1)

		resp, err := f.api.AppleGet(context.Background(), &pb.AppleGetRequest{
			Id: apple.Id,
		})

		require.NoError(t, err)
		assert.Equal(t, apple.Id, resp.GetId())
		assert.Equal(t, apple.Color.Name, resp.GetColor())
		assert.Equal(t, apple.Price, resp.GetPrice())
		assert.Equal(t, apple.Color.Id, resp.GetColorId())
	})

	t.Run("throws an error when apple is not exists", func(t *testing.T) {
		f := setUp(t)

		appleId := fake().DbId
		f.storage.EXPECT().
			Get(gomock.Any(), appleId).
			Return(nil, storage.ErrNotExists).
			Times(1)

		_, err := f.api.AppleGet(context.Background(), &pb.AppleGetRequest{
			Id: appleId,
		})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.NotFound)
	})

	t.Run("throws an internal error", func(t *testing.T) {
		f := setUp(t)

		appleId := fake().DbId
		f.storage.EXPECT().
			Get(gomock.Any(), appleId).
			Return(nil, errors.New("custom error")).
			Times(1)

		_, err := f.api.AppleGet(context.Background(), &pb.AppleGetRequest{
			Id: appleId,
		})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.Internal)
	})
}
