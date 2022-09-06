package delete

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"strconv"
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
	return "delete"
}

func (command) Description() string {
	return "<id> - delete existing v1"
}

func (c *command) Execute(args string) (string, error) {
	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "%s is not a number", args)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.stor.Delete(ctx, id); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d was was deleted", id), nil
}
