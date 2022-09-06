package client

import (
	"context"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka/message"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type client struct {
	producer kafka.Producer
	topic    string
}

func New(producer kafka.Producer, topic string) storage.AppleMutableStorage {
	return &client{
		producer: producer,
		topic:    topic,
	}
}

func (c *client) Add(ctx context.Context, entity *applePkg.Apple) error {
	return c.producer.Send(ctx, []kafka.Message{message.New(c.topic, message.CommandMessage{
		Payload: *entity,
		Command: "add",
	})})
}

func (c *client) Update(ctx context.Context, entity *applePkg.Apple) error {
	return c.producer.Send(ctx, []kafka.Message{message.New(c.topic, message.CommandMessage{
		Payload: *entity,
		Command: "update",
	})})
}

func (c *client) Delete(ctx context.Context, id uint64) error {
	return c.producer.Send(ctx, []kafka.Message{message.New(c.topic, message.CommandMessage{
		Payload: id,
		Command: "delete",
	})})
}
