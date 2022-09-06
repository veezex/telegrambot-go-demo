package item

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCacheItem(t *testing.T) {
	t.Run("can get value", func(t *testing.T) {
		setUp(t)

		val := fake().Value
		item := New(val, 1000)
		assert.Equal(t, val, item.GetValue())
	})

	t.Run("should not be expired on long living cache", func(t *testing.T) {
		setUp(t)

		item := New(fake().Value, 1*time.Hour)
		assert.False(t, item.Expired())
	})

	t.Run("can be expired", func(t *testing.T) {
		setUp(t)

		item := New(fake().Value, 0)
		assert.True(t, item.Expired())
	})
}
