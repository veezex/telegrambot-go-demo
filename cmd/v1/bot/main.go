package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/api/combined"
	grpcClient "gitlab.ozon.dev/veezex/homework/internal/pkg/api/grpc/v1/client"
	msgBrokerClient "gitlab.ozon.dev/veezex/homework/internal/pkg/api/msgbroker/client"
	producerPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/kafka/producer"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage/validation"
	settingsPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/util/settings"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
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
	runTelegramBot(validatedService, config.TelegramBotApiKey)
}
