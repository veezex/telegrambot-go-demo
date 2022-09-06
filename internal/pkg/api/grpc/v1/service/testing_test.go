package service

import (
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"log"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	mocks_storage "gitlab.ozon.dev/veezex/homework/internal/pkg/storage/mocks"
)

type Fixture struct {
	storage *mocks_storage.MockAppleStorage
	api     pb.AppleServiceServer
}

func setUp(t *testing.T) Fixture {
	// t.Parallel()

	stor := mocks_storage.NewMockAppleStorage(gomock.NewController(t))
	return Fixture{
		storage: stor,
		api:     New(stor),
	}
}

type FakeData struct {
	Price float64 `faker:"boundary_start=10, boundary_end=10000"`
	Color string  `faker:"len=25"`
	DbId  uint64  `faker:"boundary_start=1, boundary_end=1000"`
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
