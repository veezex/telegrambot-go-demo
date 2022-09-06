package monitoring

import (
	"expvar"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/locker"
	"strconv"
)

var cacheHit Counter
var cacheMiss Counter
var inRequests Counter
var successRequests Counter
var failedRequests Counter

type Counter interface {
	Inc()
	String() string
}

type counter struct {
	cnt    int64
	locker locker.Locker
}

func NewCounter() Counter {
	return &counter{
		locker: locker.New(),
	}
}

func (c *counter) Inc() {
	defer c.locker.Lock()()
	c.cnt++
}

func (c *counter) String() string {
	defer c.locker.RLock()()
	return strconv.FormatInt(c.cnt, 10)
}

func init() {
	cacheMiss = NewCounter()
	expvar.Publish("cacheMiss", cacheMiss)

	cacheHit = NewCounter()
	expvar.Publish("cacheHit", cacheHit)

	inRequests = NewCounter()
	expvar.Publish("inRequests", inRequests)

	successRequests = NewCounter()
	expvar.Publish("successRequests", successRequests)

	failedRequests = NewCounter()
	expvar.Publish("failedRequests", failedRequests)
}
