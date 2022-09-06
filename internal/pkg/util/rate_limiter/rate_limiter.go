package rate_limiter

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrTimeout = errors.New("Timeout exceeded")
)

type RateLimiter interface {
	Run(handler) (interface{}, error)
}

type handler func() (interface{}, error)

type limiter struct {
	timeout time.Duration
	poolCh  chan struct{}
}

func New(maxWorkers uint, timeout time.Duration) RateLimiter {
	return &limiter{
		poolCh:  make(chan struct{}, maxWorkers),
		timeout: timeout,
	}
}

func (l *limiter) Run(job handler) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), l.timeout)
	defer cancel()

	select {
	case l.poolCh <- struct{}{}:
		result, err := job()
		<-l.poolCh
		return result, err
	case <-ctx.Done():
		return nil, ErrTimeout
	}
}
