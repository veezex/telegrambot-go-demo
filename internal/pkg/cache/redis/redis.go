package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/cache"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/monitoring"
	"time"
)

type redisCache struct {
	rdb     *redis.Client
	timeout time.Duration
	ctx     context.Context
	monitor monitoring.CacheMonitor
}

func New(ctx context.Context, timeout time.Duration, rdb *redis.Client) cache.Cacher {
	return &redisCache{
		ctx:     ctx,
		rdb:     rdb,
		timeout: timeout,
	}
}

func (rc *redisCache) SetMonitor(monitor monitoring.CacheMonitor) cache.Cacher {
	rc.monitor = monitor
	return rc
}

func (rc *redisCache) Get(key string, getter func() (cache.CacheValue, error)) (cache.CacheValue, error) {
	val, err := rc.rdb.Get(rc.ctx, key).Bytes()
	if err == nil {
		if rc.monitor != nil {
			rc.monitor.Hit()
		}

		return val, nil
	}

	if rc.monitor != nil {
		rc.monitor.Miss()
	}

	newVal, err := getter()
	if err != nil {
		return nil, err
	}

	if err := rc.rdb.Set(rc.ctx, key, newVal, rc.timeout).Err(); err != nil {
		return nil, err
	}

	return newVal, nil
}

func (rc *redisCache) Invalidate(key string) error {
	return rc.rdb.Del(rc.ctx, key).Err()
}
