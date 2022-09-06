package service

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc"
)

type streamListMock struct {
	grpc.ServerStream
	ctx    context.Context
	sentCh chan *pb.AppleGetResponse
}

func newStreamListMock() *streamListMock {
	return &streamListMock{
		ctx:    context.Background(),
		sentCh: make(chan *pb.AppleGetResponse, 1),
	}
}

func (m *streamListMock) extractList() []applePkg.Apple {
	list := make([]applePkg.Apple, 0)

	done := make(chan struct{})
	go func() {
		for {
			select {
			case resp := <-m.sentCh:
				list = append(list, applePkg.Apple{
					Id: resp.GetId(),
					Color: colorPkg.Color{
						Id:   resp.GetColorId(),
						Name: resp.GetColor(),
					},
					Price: resp.GetPrice(),
				})
			default:
				done <- struct{}{}
				return
			}
		}
	}()
	<-done

	return list
}

func (m *streamListMock) Send(resp *pb.AppleGetResponse) error {
	m.sentCh <- resp
	return nil
}

func (m *streamListMock) Context() context.Context {
	return m.ctx
}
