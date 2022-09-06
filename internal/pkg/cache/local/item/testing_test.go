package item

import (
	"github.com/bxcodec/faker/v4"
	"log"
	"testing"
)

func setUp(t *testing.T) {
	t.Parallel()
}

type FakeData struct {
	Value string `faker:"sentence"`
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
