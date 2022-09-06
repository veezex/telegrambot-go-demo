package entities

import "github.com/pkg/errors"

var (
	ErrParseUndefined = errors.New("Parse error, undefined field")
	ErrParseType      = errors.New("Parse error, wrong type")
)
