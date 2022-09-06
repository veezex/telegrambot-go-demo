package router

import (
	"fmt"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"strings"

	"github.com/pkg/errors"
)

var (
	errCommandExists  = errors.New("Command already exists")
	errCommandUnknown = errors.New("Unknown command")
)

type router struct {
	commands map[string]cmdPkg.Command
}

func New() Router {
	r := &router{}
	r.commands = make(map[string]cmdPkg.Command)

	return r
}

func (r *router) RegisterCommand(cmd cmdPkg.Command) error {
	name := cmd.Name()
	if _, ok := r.commands[name]; ok {
		return errCommandExists
	}
	r.commands[name] = cmd
	return nil
}

func (r *router) HandleCommand(cmd, args string) (string, error) {
	if command, ok := r.commands[cmd]; ok {
		return command.Execute(args)
	} else {
		return "", errCommandUnknown
	}
}

func (r *router) String() string {
	hints := make([]string, 0, len(r.commands))

	for _, v := range r.commands {
		hints = append(hints, fmt.Sprintf("/%s - %s", v.Name(), v.Description()))
	}

	return strings.Join(hints, "\n")
}
