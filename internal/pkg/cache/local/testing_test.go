package local

import (
	"github.com/bxcodec/faker/v4"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/cache"
	"log"
	"testing"
	"time"
)

type Fixture struct {
	cache cache.Cacher
}

func setUp(t *testing.T) Fixture {
	//t.Parallel()

	return Fixture{
		cache: New(1 * time.Hour),
	}
}

type FakeData struct {
	Key   string `faker:"len=25"`
	Value string `faker:"sentence"`
	Int   int
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
