package item

import (
	"gitlab.ozon.dev/veezex/homework/internal/pkg/cache"
	"time"
)

type cacheItem struct {
	value   cache.CacheValue
	expires time.Time
}

func New(value cache.CacheValue, timeout time.Duration) cache.CacheItem {
	return &cacheItem{
		expires: time.Now().Add(timeout),
		value:   value,
	}
}

func (ci *cacheItem) Expired() bool {
	return ci.expires.Before(time.Now())
}

func (ci *cacheItem) GetValue() cache.CacheValue {
	return ci.value
}
