package update

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
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
	return "update"
}

func (command) Description() string {
	return "<id> <color> <price> - update existing v1"
}

func (c *command) Execute(args string) (string, error) {
	params := strings.Split(args, " ")
	if len(params) != 3 {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "%d items: <%v>", len(params), params)
	}

	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "%s is not a number", args)
	}

	price, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return "", errors.Wrapf(cmdPkg.ErrBadArgument, "incorrect price <%s>", params[2])
	}

	a := applePkg.Apple{
		Color: colorPkg.Color{
			Name: params[1],
		},
		Price: price,
		Id:    id,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.stor.Update(ctx, &a); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s was updated", a.String()), nil
}
