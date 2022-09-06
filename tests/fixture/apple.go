//go:build integration
// +build integration

package fixture

import "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"

type appleBuilder struct {
	instance *apple.Apple
}

func Apple() *appleBuilder {
	return &appleBuilder{
		instance: &apple.Apple{},
	}
}

func (b *appleBuilder) Id(id uint64) *appleBuilder {
	b.instance.Id = id
	return b
}

func (b *appleBuilder) ColorId(id uint64) *appleBuilder {
	b.instance.Color.Id = id
	return b
}

func (b *appleBuilder) ColorName(v string) *appleBuilder {
	b.instance.Color.Name = v
	return b
}

func (b *appleBuilder) Price(v float64) *appleBuilder {
	b.instance.Price = v
	return b
}
func (b *appleBuilder) P() *apple.Apple {
	return b.instance
}

func (b *appleBuilder) V() apple.Apple {
	return *b.instance
}
