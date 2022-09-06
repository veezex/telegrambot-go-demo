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
)

func TestDelete(t *testing.T) {
	t.Run("can delete an apple", func(t *testing.T) {
		f := setUp(t)

		appleId := fake().DbId
		f.storage.EXPECT().
			Delete(gomock.Any(), appleId).
			Return(nil).
			Times(1)

		_, err := f.api.AppleDelete(context.Background(), &pb.AppleDeleteRequest{
			Id: appleId,
		})

		require.NoError(t, err)
	})

	t.Run("throws an error when apple is not exists", func(t *testing.T) {
		f := setUp(t)

		appleId := fake().DbId
		f.storage.EXPECT().
			Delete(gomock.Any(), appleId).
			Return(storage.ErrNotExists).
			Times(1)

		_, err := f.api.AppleDelete(context.Background(), &pb.AppleDeleteRequest{
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
			Delete(gomock.Any(), appleId).
			Return(errors.New("custom erros")).
			Times(1)

		_, err := f.api.AppleDelete(context.Background(), &pb.AppleDeleteRequest{
			Id: appleId,
		})

		require.NotNil(t, err)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.Internal)
	})
}
