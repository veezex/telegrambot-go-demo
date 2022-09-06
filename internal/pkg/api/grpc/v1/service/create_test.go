package service

import (
	"context"
	"errors"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	t.Run("can create an apple", func(t *testing.T) {
		f := setUp(t)

		apple := applePkg.Apple{
			Color: color.Color{
				Name: fake().Color,
			},
			Price: fake().Price,
		}

		f.storage.EXPECT().
			Add(gomock.Any(), &apple).
			Return(nil).
			Times(1)

		resp, err := f.api.AppleCreate(context.Background(), &pb.AppleCreateRequest{
			Color: apple.Color.Name,
			Price: apple.Price,
		})

		require.NoError(t, err)
		assert.Equal(t, resp.GetColorId(), uint64(0))
		assert.Equal(t, resp.GetId(), uint64(0))
	})

	t.Run("throws an error when apple is already exists", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().
			Add(gomock.Any(), gomock.Any()).
			Return(storage.ErrExists).
			Times(1)

		_, err := f.api.AppleCreate(context.Background(), &pb.AppleCreateRequest{})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.AlreadyExists)
	})

	t.Run("throws an internal error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().
			Add(gomock.Any(), gomock.Any()).
			Return(errors.New("Custom error")).
			Times(1)

		_, err := f.api.AppleCreate(context.Background(), &pb.AppleCreateRequest{})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.Internal)
	})
}
