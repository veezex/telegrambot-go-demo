package color

import (
	"github.com/bxcodec/faker/v4"
	"log"
	"testing"
)

type Fixture struct {
	color Color
}

func setUp(t *testing.T) Fixture {
	//t.Parallel()
	return Fixture{
		color: Color{
			Id:   fake().DbId,
			Name: fake().Color,
		},
	}
}

type FakeData struct {
	DbId      uint64 `faker:"boundary_start=1, boundary_end=1000"`
	Color     string `faker:"len=25"`
	LongColor string `faker:"len=125"`
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
