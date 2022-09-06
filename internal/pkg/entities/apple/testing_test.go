package apple

import (
	"github.com/bxcodec/faker/v4"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	"log"
	"testing"
)

type Fixture struct {
	apple Apple
}

func setUp(t *testing.T) Fixture {
	//t.Parallel()
	return Fixture{
		apple: Apple{
			Id: fake().DbId,
			Color: colorPkg.Color{
				Id:   fake().DbId,
				Name: fake().Color,
			},
			Price: fake().Price,
		},
	}
}

type FakeData struct {
	DbId  uint64  `faker:"boundary_start=1, boundary_end=1000"`
	Price float64 `faker:"boundary_start=10, boundary_end=10000"`
	Color string  `faker:"len=25"`
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
