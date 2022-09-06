package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	apiPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/api/grpc/v2/service"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
	"google.golang.org/grpc"
	"net"
	"time"
)

func runGRPCServer(s storage.AppleStorage, port uint64, rdb *redis.Client, pubSubChannel string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAppleServiceServer(grpcServer, apiPkg.New(s, func(ctx context.Context, value []byte, meta interface{}) error {
		key := uuid.NewString()

		// publish result in redis cache
		err := rdb.HSet(ctx, key, "value", value, "meta", meta).Err()
		if err != nil {
			return err
		}

		err = rdb.Expire(ctx, key, 10*time.Second).Err()
		if err != nil {
			return err
		}

		err = rdb.Publish(ctx, pubSubChannel, key).Err()
		if err != nil {
			return err
		}

		return nil
	}))

	if err = grpcServer.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
