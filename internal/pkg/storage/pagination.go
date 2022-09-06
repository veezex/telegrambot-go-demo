package storage

import (
	"github.com/pkg/errors"
	"strings"
)

type PaginationOpts struct {
	Order  string
	Limit  uint64
	Offset uint64
}

var (
	ErrOptionInvalid = errors.New("Invalid option")
)

func (po *PaginationOpts) Validate() error {
	if po.Limit <= 0 {
		return errors.Wrapf(ErrOptionInvalid, "Limit should be more than 0, <%d>", po.Limit)
	}

	if po.Offset < 0 {
		return errors.Wrapf(ErrOptionInvalid, "Limit should be positive, <%d>", po.Offset)
	}

	if strings.ToUpper(po.Order) != "ASC" && strings.ToUpper(po.Order) != "DESC" {
		return errors.Wrapf(ErrOptionInvalid, "Order should be either asc or desc, <%s>", po.Order)
	}

	return nil
}
