package apple

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
)

type Apple struct {
	Id    uint64         `db:"id"`
	Color colorPkg.Color `db:"color_id"`
	Price float64        `db:"price"`
}

func ParseFromMap(in map[string]interface{}) (*Apple, error) {
	// get id
	idRaw, okIdRaw := in["Id"]
	if !okIdRaw {
		return nil, errors.Wrap(entities.ErrParseUndefined, "apple id")
	}

	id, okId := idRaw.(float64)
	if !okId {
		return nil, errors.Wrapf(entities.ErrParseType, "wrong apple id type <%t>", idRaw)
	}

	// get price
	priceRaw, okPriceRaw := in["Price"]
	if !okPriceRaw {
		return nil, errors.Wrap(entities.ErrParseUndefined, "apple price")
	}

	price, okPrice := priceRaw.(float64)
	if !okPrice {
		return nil, errors.Wrapf(entities.ErrParseType, "wrong apple price type <%t>", priceRaw)
	}

	// get color
	colorRaw, okColorRaw := in["Color"]
	if !okColorRaw {
		return nil, errors.Wrap(entities.ErrParseUndefined, "apple color")
	}

	color, okColor := colorRaw.(map[string]interface{})
	if !okColor {
		return nil, errors.Wrapf(entities.ErrParseType, "wrong apple color type <%t>", colorRaw)
	}

	colorParsed, err := colorPkg.ParseFromMap(color)
	if err != nil {
		return nil, err
	}

	return &Apple{
		Id:    uint64(id),
		Price: price,
		Color: *colorParsed,
	}, nil
}

func Validate(a *Apple) error {
	if a.Price < 0 {
		return errors.Wrapf(entities.ErrValidation, "Price can not be lower than 0, <%f>", a.Price)
	}

	if err := colorPkg.Validate(&a.Color); err != nil {
		return err
	}

	return nil
}

func (a *Apple) String() string {
	if a.Id == 0 {
		return fmt.Sprintf("Apple: color %s / price %f", a.Color.String(), a.Price)
	}
	return fmt.Sprintf("Apple %d: color %s / price %f", a.Id, a.Color.String(), a.Price)
}
