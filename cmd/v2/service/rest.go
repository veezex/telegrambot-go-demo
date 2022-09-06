package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/pkg/api"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func runRESTServer(grpcPort uint64, restPort uint64) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAppleServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf(":%d", grpcPort), opts); err != nil {
		logrus.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, _ *http.Request) {
		w.Write(api.SwaggerJson)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", restPort), mux); err != nil {
		logrus.Fatal(err)
	}
}
