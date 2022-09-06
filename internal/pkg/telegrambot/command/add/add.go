package add

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"strconv"
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
	return "add"
}

func (command) Description() string {
	return "<color> <price> - add a new v1"
}

func (c *command) Execute(args string) (string, error) {
	params := strings.Split(args, " ")
	if len(params) != 2 {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "%d items: <%v>", len(params), params)
	}

	price, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "incorrect price <%s>", params[1])
	}

	a := apple.Apple{
		Color: color.Color{
			Name: params[0],
		},
		Price: price,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = c.stor.Add(ctx, &a)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s was added", a.String()), nil
}
