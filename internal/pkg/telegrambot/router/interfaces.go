package router

import (
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
)

type CommandExecutor interface {
	HandleCommand(cmd, args string) (string, error)
}

type CommandRegistrator interface {
	RegisterCommand(cmd cmdPkg.Command) error
}

type Router interface {
	CommandExecutor
	CommandRegistrator
	String() string
}
