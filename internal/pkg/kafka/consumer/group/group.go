package group

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka"
	"time"
)

type group struct {
	cg sarama.ConsumerGroup
}

func New(groupId string, brokers []string) (kafka.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		return nil, err
	}

	return &group{
		cg: client,
	}, nil
}

func (g *group) Listen(ctx context.Context, topics []string, consumer sarama.ConsumerGroupHandler) error {
	for {
		if err := g.cg.Consume(ctx, topics, consumer); err != nil {
			logrus.Printf("on consume: %v", err)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
