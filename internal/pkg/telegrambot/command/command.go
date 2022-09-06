//go:generate mockgen -source=./command.go -destination=./mocks/command.go -package=mocks
package command

import "github.com/pkg/errors"

type Command interface {
	Execute(args string) (string, error)
	Name() string
	Description() string
}

var ErrBadArgument = errors.New("Bad argument")
