package router

import (
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
)

type command struct {
	router Router
}

func NewHelpCommand(r Router) cmdPkg.Command {
	return &command{
		router: r,
	}
}

func (command) Name() string {
	return "help"
}

func (command) Description() string {
	return "list of commands"
}

func (c *command) Execute(_ string) (string, error) {
	return c.router.String(), nil
}
