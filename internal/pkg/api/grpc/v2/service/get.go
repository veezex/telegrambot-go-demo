package service

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v2"
	"golang.org/x/net/context"
)

func (i *impl) AppleGet(_ context.Context, in *pb.AppleGetRequest) (*pb.AppleGetResponse, error) {
	inId := in.GetId()
	go func() {
		bgCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		apple, err := i.stor.Get(bgCtx, inId)
		if err != nil {
			logrus.Error(err)
			return
		}

		appleStr, err := json.Marshal(apple)
		if err != nil {
			logrus.Error(err)
			return
		}

		err = i.publishResult(bgCtx, appleStr, "get")
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	return &pb.AppleGetResponse{}, nil
}
