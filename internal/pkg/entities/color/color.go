package color

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities"
)

type Color struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}

func ParseFromMap(in map[string]interface{}) (*Color, error) {
	// get id
	idRaw, okIdRaw := in["Id"]
	if !okIdRaw {
		return nil, errors.Wrap(entities.ErrParseUndefined, "color id")
	}

	id, okId := idRaw.(float64)
	if !okId {
		return nil, errors.Wrapf(entities.ErrParseType, "wrong color id type <%t>", idRaw)
	}

	// get name
	nameRaw, okNameRaw := in["Name"]
	if !okNameRaw {
		return nil, errors.Wrap(entities.ErrParseUndefined, "color name")
	}

	name, okName := nameRaw.(string)
	if !okName {
		return nil, errors.Wrapf(entities.ErrParseType, "wrong color name type <%t>", idRaw)
	}

	return &Color{
		Id:   uint64(id),
		Name: name,
	}, nil
}

func Validate(c *Color) error {
	if len(c.Name) > 100 {
		return errors.Wrapf(entities.ErrValidation, "Color name can not be longer than 100 characters, <%s>", c.Name)
	}

	return nil
}

func (c *Color) String() string {
	if c.Id == 0 {
		return fmt.Sprintf("Color: color %s", c.Name)
	}
	return fmt.Sprintf("Color %d: color %s", c.Id, c.Name)
}
