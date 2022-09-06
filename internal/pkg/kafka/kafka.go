package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

type Producer interface {
	Send(ctx context.Context, msg []Message) error
}

type ConsumerGroup interface {
	Listen(ctx context.Context, topics []string, consumer sarama.ConsumerGroupHandler) error
}

type Message interface {
	SetKey(string) Message
	SaramaMessage() (*sarama.ProducerMessage, error)
}

type Router interface {
	ExecuteCommand(ctx context.Context, payload []byte) error
}
