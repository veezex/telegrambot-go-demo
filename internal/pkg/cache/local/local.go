package local

import (
	"gitlab.ozon.dev/veezex/homework/internal/pkg/cache"
	itemPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/cache/local/item"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/monitoring"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/locker"
	"time"
)

type localCache struct {
	locker  locker.Locker
	bucket  map[string]cache.CacheItem
	timeout time.Duration
	monitor monitoring.CacheMonitor
}

func New(timeout time.Duration) cache.Cacher {
	return &localCache{
		locker:  locker.New(),
		bucket:  make(map[string]cache.CacheItem),
		timeout: timeout,
	}
}

func (lc *localCache) SetMonitor(monitor monitoring.CacheMonitor) cache.Cacher {
	lc.monitor = monitor
	return lc
}

func (lc *localCache) Get(key string, getter func() (cache.CacheValue, error)) (cache.CacheValue, error) {
	item := lc.getItem(key)

	if item != nil && !item.Expired() {
		if lc.monitor != nil {
			lc.monitor.Hit()
		}

		return item.GetValue(), nil
	}

	if lc.monitor != nil {
		lc.monitor.Miss()
	}

	newVal, err := getter()
	if err != nil {
		return nil, err
	}
	lc.setItem(key, itemPkg.New(newVal, lc.timeout))

	return newVal, nil
}

func (lc *localCache) Invalidate(key string) error {
	defer lc.locker.Lock()()

	if _, ok := lc.bucket[key]; ok {
		delete(lc.bucket, key)
	}

	return nil
}

func (lc *localCache) getItem(key string) cache.CacheItem {
	defer lc.locker.RLock()()

	if _, ok := lc.bucket[key]; ok {
		return lc.bucket[key]
	}

	return nil
}

func (lc *localCache) setItem(key string, item cache.CacheItem) {
	defer lc.locker.Lock()()
	lc.bucket[key] = item
}
