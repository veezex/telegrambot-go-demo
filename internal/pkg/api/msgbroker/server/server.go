package server

import (
	"context"
	"errors"
	"fmt"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/kafka/message"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

var (
	errUnknownPayload = errors.New("Unknown payload")
)

type commandHandler = func(context.Context, storage.AppleMutableStorage, interface{}) error

type server struct {
	storage  storage.AppleMutableStorage
	handlers map[string]commandHandler
}

func New(storage storage.AppleMutableStorage) kafka.Router {
	result := &server{
		storage:  storage,
		handlers: make(map[string]commandHandler),
	}

	result.registerCommand("add", addCmd)
	result.registerCommand("update", updateCmd)
	result.registerCommand("delete", deleteCmd)

	return result
}

func (s *server) registerCommand(command string, handler commandHandler) {
	s.handlers[command] = handler
}

func (s *server) ExecuteCommand(ctx context.Context, in []byte) error {
	msg, err := message.ParseCommandMessage(in)

	if err != nil {
		return err
	}

	if handler, ok := s.handlers[msg.Command]; ok {
		return handler(ctx, s.storage, msg.Payload)
	}

	return errors.New(fmt.Sprintf("Command is not exists <%v>", msg.Command))
}
