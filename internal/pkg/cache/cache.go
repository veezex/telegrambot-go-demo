//go:generate mockgen -source=./cache.go -destination=./mocks/cache.go -package=mocks_cache
package cache

import "gitlab.ozon.dev/veezex/homework/internal/pkg/monitoring"

type CacheValue = []byte

type Cacher interface {
	Get(key string, getter func() (CacheValue, error)) (CacheValue, error)
	Invalidate(key string) error
	SetMonitor(monitor monitoring.CacheMonitor) Cacher
}

type CacheItem interface {
	Expired() bool
	GetValue() CacheValue
}
