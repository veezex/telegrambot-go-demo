package monitoring

type CacheMonitor interface {
	Hit()
	Miss()
}

type cacheMonitoring struct{}

func NewCacheMonitoring() CacheMonitor {
	return &cacheMonitoring{}
}

func (cacheMonitoring) Hit() {
	cacheHit.Inc()
}

func (cacheMonitoring) Miss() {
	cacheMiss.Inc()
}
