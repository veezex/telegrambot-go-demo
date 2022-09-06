package message

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka"
)

type message struct {
	key     *string
	topic   string
	payload interface{}
}

func New(topic string, payload interface{}) kafka.Message {
	return &message{
		payload: payload,
		topic:   topic,
	}
}

func (m *message) SetKey(key string) kafka.Message {
	m.key = &key
	return m
}

func (m *message) SaramaMessage() (*sarama.ProducerMessage, error) {
	var key sarama.StringEncoder
	if m.key != nil {
		key = sarama.StringEncoder(*m.key)
	}

	value, err := json.Marshal(m.payload)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic: m.topic,
		Key:   key,
		Value: sarama.ByteEncoder(value),
	}, nil
}
