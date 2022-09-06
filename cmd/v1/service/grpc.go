package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	apiPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/api/grpc/v1/service"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc"
	"net"
)

func runGRPCServer(s storage.AppleStorage, port uint64) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAppleServiceServer(grpcServer, apiPkg.New(s))

	if err = grpcServer.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
