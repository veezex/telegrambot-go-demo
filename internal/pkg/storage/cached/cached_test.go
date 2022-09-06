package cached

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCached_Add(t *testing.T) {
	t.Run("can trigger add action", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().
			Add(gomock.Any(), &f.apple).Return(nil).Times(1)
		f.cacher.EXPECT().
			Invalidate(makeKey(f.apple.Id)).Return(nil).Times(1)

		err := f.impl.Add(context.Background(), &f.apple)

		assert.NoError(t, err)
	})

	t.Run("throws error on covered storage error", func(t *testing.T) {
		f := setUp(t)
		errThrown := errors.New("storage error")

		f.storage.EXPECT().
			Add(gomock.Any(), &f.apple).Return(errThrown).Times(1)

		err := f.impl.Add(context.Background(), &f.apple)

		assert.ErrorIs(t, err, errThrown)
	})

	t.Run("throws error on cache error", func(t *testing.T) {
		f := setUp(t)
		errThrown := errors.New("storage error")

		f.storage.EXPECT().
			Add(gomock.Any(), &f.apple).Return(nil).Times(1)
		f.cacher.EXPECT().
			Invalidate(makeKey(f.apple.Id)).Return(errThrown).Times(1)

		err := f.impl.Add(context.Background(), &f.apple)

		assert.ErrorIs(t, err, errThrown)
	})
}

func TestCached_Delete(t *testing.T) {
	t.Run("can trigger delete action", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().
			Delete(gomock.Any(), f.apple.Id).Return(nil).Times(1)
		f.cacher.EXPECT().
			Invalidate(makeKey(f.apple.Id)).Return(nil).Times(1)

		err := f.impl.Delete(context.Background(), f.apple.Id)

		assert.NoError(t, err)
	})

	t.Run("throws error on covered storage error", func(t *testing.T) {
		f := setUp(t)
		errThrown := errors.New("storage error")

		f.storage.EXPECT().
			Delete(gomock.Any(), f.apple.Id).Return(errThrown).Times(1)

		err := f.impl.Delete(context.Background(), f.apple.Id)

		assert.ErrorIs(t, err, errThrown)
	})

	t.Run("throws error on cache storage error", func(t *testing.T) {
		f := setUp(t)
		errThrown := errors.New("storage error")

		f.storage.EXPECT().
			Delete(gomock.Any(), f.apple.Id).Return(nil).Times(1)
		f.cacher.EXPECT().
			Invalidate(makeKey(f.apple.Id)).Return(errThrown).Times(1)

		err := f.impl.Delete(context.Background(), f.apple.Id)

		assert.ErrorIs(t, err, errThrown)
	})
}

func TestCached_Get(t *testing.T) {
	t.Run("can get cached item", func(t *testing.T) {
		f := setUp(t)

		f.cacher.EXPECT().
			Get(makeKey(f.apple.Id), gomock.Any()).
			Return(&getResult{value: &f.apple, err: nil}, nil).Times(1)

		apple, err := f.impl.Get(context.Background(), f.apple.Id)

		require.NoError(t, err)
		assert.Equal(t, f.apple, *apple)
	})

	t.Run("throws an error when cache throws an error", func(t *testing.T) {
		f := setUp(t)
		errThrown := errors.New("storage error")

		f.cacher.EXPECT().
			Get(makeKey(f.apple.Id), gomock.Any()).
			Return(&getResult{value: nil, err: errThrown}, nil).Times(1)

		_, err := f.impl.Get(context.Background(), f.apple.Id)
		assert.ErrorIs(t, err, errThrown)
	})

	t.Run("throws an error when cache result is malformated", func(t *testing.T) {
		f := setUp(t)

		f.cacher.EXPECT().
			Get(makeKey(f.apple.Id), gomock.Any()).
			Return(1, nil).Times(1)

		_, err := f.impl.Get(context.Background(), f.apple.Id)
		assert.ErrorIs(t, err, errMalformatted)
	})

	t.Run("throws an error when not apple is cached", func(t *testing.T) {
		f := setUp(t)

		f.cacher.EXPECT().
			Get(makeKey(f.apple.Id), gomock.Any()).
			Return(&getResult{value: 1, err: nil}, nil).Times(1)

		_, err := f.impl.Get(context.Background(), f.apple.Id)
		assert.ErrorIs(t, err, errMalformatted)
	})
}

func TestCached_List(t *testing.T) {
	t.Run("can get list of cached items", func(t *testing.T) {
		f := setUp(t)

		f.cacher.EXPECT().
			Get(makeKey(f.apple.Id), gomock.Any()).
			Return(&getResult{value: &f.apple, err: nil}, nil).Times(1)

		apple, err := f.impl.Get(context.Background(), f.apple.Id)

		require.NoError(t, err)
		assert.Equal(t, f.apple, *apple)
	})
}
