package local

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/cache"
	"testing"
)

func TestCache(t *testing.T) {
	t.Run("can get cached value", func(t *testing.T) {
		f := setUp(t)

		val := fake().Value
		key := fake().Key

		retVal, err := f.cache.Get(key, func() (cache.CacheValue, error) {
			return cache.CacheValue(val), nil
		})

		require.NoError(t, err)
		assert.Equal(t, cache.CacheValue(val), retVal)

		retVal, err = f.cache.Get(key, func() (cache.CacheValue, error) {
			return cache.CacheValue(val + "wrong_value"), nil
		})

		require.NoError(t, err)
		assert.Equal(t, cache.CacheValue(val), retVal)
	})

	t.Run("can invalidate cache", func(t *testing.T) {
		f := setUp(t)

		val := fake().Value
		key := fake().Key

		retVal, err := f.cache.Get(key, func() (cache.CacheValue, error) {
			return cache.CacheValue(val), nil
		})

		require.NoError(t, err)
		assert.Equal(t, cache.CacheValue(val), retVal)

		err = f.cache.Invalidate(key)
		require.NoError(t, err)

		retVal, err = f.cache.Get(key, func() (cache.CacheValue, error) {
			return cache.CacheValue(val + "wrong_value"), nil
		})

		require.NoError(t, err)
		assert.Equal(t, cache.CacheValue(val+"wrong_value"), retVal)
	})
}
