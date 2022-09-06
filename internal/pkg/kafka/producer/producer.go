package producer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka"
)

type producer struct {
	producer sarama.SyncProducer
}

func New(brokers []string) (kafka.Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "producer creation error")
	}

	return &producer{
		producer: syncProducer,
	}, nil
}

func (p *producer) Send(_ context.Context, msgs []kafka.Message) error {
	outMsgs := make([]*sarama.ProducerMessage, len(msgs))
	for index, m := range msgs {
		outMsg, err := m.SaramaMessage()
		if err != nil {
			return err
		}

		outMsgs[index] = outMsg
	}

	return p.producer.SendMessages(outMsgs)
}
