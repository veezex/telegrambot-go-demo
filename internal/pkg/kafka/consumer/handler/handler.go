package handler

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka"
)

type handler struct {
	router kafka.Router
}

func New(router kafka.Router) sarama.ConsumerGroupHandler {
	return &handler{
		router: router,
	}
}

func (h *handler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *handler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case <-session.Context().Done():
			logrus.Println("consumer: done")
			return nil
		case msg := <-claim.Messages():
			logrus.Println("%v", string(msg.Value))

			if err := h.router.ExecuteCommand(session.Context(), msg.Value); err != nil {
				logrus.Errorf("%v", err)
				return err
			}

			session.MarkMessage(msg, "")
		}
	}
}
