package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	msgBrokerServer "gitlab.ozon.dev/veezex/homework/internal/pkg/api/msgbroker/server"
	redisCache "gitlab.ozon.dev/veezex/homework/internal/pkg/cache/redis"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka/consumer/group"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka/consumer/handler"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/monitoring"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage/cached"
	postgressPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage/postgress"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage/throttled"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/postgress"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/rate_limiter"
	settingsPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/util/settings"
	"time"
)

func main() {
	config := settingsPkg.New(".env").GetConfig()

	// init postgress connection
	p := &postgress.Postgress{
		Host:     config.DbHost,
		Port:     config.DbPort,
		User:     config.DbUser,
		Password: config.DbPassword,
		DBname:   config.DbName,
	}
	pool, closeConnection := p.CreatePool()
	defer closeConnection()

	// create pesistent storage instance (postgress)
	postgressStorage := postgressPkg.New(pool)

	// wrap storage with a cache
	//cachedStorage := cached.New(
	//	postgressStorage,
	//	local.New(time.Second*10).SetMonitor(monitoring.NewCacheMonitoring()),
	//)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       int(config.RedisDb),
	})
	cachedStorage := cached.New(
		postgressStorage,
		redisCache.New(context.Background(), time.Second*50, rdb).
			SetMonitor(monitoring.NewCacheMonitoring()),
	)

	// add call throttling to a cached storage
	throttledStorage := throttled.New(cachedStorage, rate_limiter.New(10, time.Second*10))

	// add api monitoring
	monitoredStorage := monitoring.NewStorageMonitoring(throttledStorage)

	// run debug webserver
	go debugCounters(config.DebugPort)

	// kafka consumer
	consumerGroup, err := group.New(
		config.KafkaGroupId,
		config.KafkaBrokers)
	if err != nil {
		logrus.Fatal(err)
	}

	go func() {
		err := consumerGroup.Listen(
			context.Background(),
			[]string{config.KafkaTopic},
			handler.New(msgBrokerServer.New(monitoredStorage)),
		)
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	// grpc & rest server
	go runGRPCServer(monitoredStorage, config.GrpcServerPort)
	runRESTServer(config.GrpcServerPort, config.RestServerPort)
}
