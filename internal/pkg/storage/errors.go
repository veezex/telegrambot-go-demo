package storage

import "github.com/pkg/errors"

var (
	ErrNotExists = errors.New("Item does not exists")
	ErrExists    = errors.New("Item already exists")
)
