package get

import (
	"context"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"strconv"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type command struct {
	stor storage.AppleStorage
}

func New(s storage.AppleStorage) cmdPkg.Command {
	return &command{
		stor: s,
	}
}

func (command) Name() string {
	return "get"
}

func (command) Description() string {
	return "<id> - get existing v1"
}

func (c *command) Execute(args string) (string, error) {
	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "%s is not a number", args)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	entity, err := c.stor.Get(ctx, id)
	if err != nil {
		return "", err
	}

	if entity == nil {
		return "", nil
	}

	return entity.String(), nil
}
