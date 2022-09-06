package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
)

func TestUpdate(t *testing.T) {
	t.Run("can update an apple", func(t *testing.T) {
		f := setUp(t)

		apple := applePkg.Apple{
			Id: fake().DbId,
			Color: color.Color{
				Name: fake().Color,
			},
			Price: fake().Price,
		}

		f.storage.EXPECT().
			Update(gomock.Any(), &apple).
			Return(nil).
			Times(1)

		_, err := f.api.AppleUpdate(context.Background(), &pb.AppleUpdateRequest{
			Id:    apple.Id,
			Color: apple.Color.Name,
			Price: apple.Price,
		})

		require.NoError(t, err)
	})

	t.Run("throws an error when apple is not exists", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().
			Update(gomock.Any(), gomock.Any()).
			Return(storage.ErrNotExists).
			Times(1)

		_, err := f.api.AppleUpdate(context.Background(), &pb.AppleUpdateRequest{})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.NotFound)
	})

	t.Run("throws an internal error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().
			Update(gomock.Any(), gomock.Any()).
			Return(errors.New("custom error")).
			Times(1)

		_, err := f.api.AppleUpdate(context.Background(), &pb.AppleUpdateRequest{})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.Internal)
	})
}
