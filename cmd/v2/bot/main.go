package main

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/api/combined"
	grpcClient "gitlab.ozon.dev/veezex/homework/internal/pkg/api/grpc/v2/client"
	msgBrokerClient "gitlab.ozon.dev/veezex/homework/internal/pkg/api/msgbroker/client"
	producerPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/kafka/producer"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage/validation"
	settingsPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/util/settings"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := settingsPkg.New(".env").GetConfig()

	// get grpc client
	storage, err := grpc.Dial(fmt.Sprintf(":%d", config.GrpcServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}

	// remote storage
	producer, err := producerPkg.New(config.KafkaBrokers)
	if err != nil {
		logrus.Fatal(err)
	}
	service := combined.New(
		msgBrokerClient.New(producer, config.KafkaTopic),
		grpcClient.New(pb.NewAppleServiceClient(storage)),
	)

	// adding validation layer
	validatedService := validation.New(service)

	// telegrambot server
	runTelegramBot(validatedService, config.TelegramBotApiKey, redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       int(config.RedisDb),
	}), config.RedisPubSubChannel)
}
