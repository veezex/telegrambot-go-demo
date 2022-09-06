package list

import (
	"context"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"strings"
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
	return "list"
}

func (command) Description() string {
	return "list data"
}

func (c *command) Execute(_ string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := c.stor.List(ctx, nil)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", nil
	}

	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}

	if len(res) == 0 {
		return "List is empty", nil
	}

	return strings.Join(res, "\n"), nil
}
