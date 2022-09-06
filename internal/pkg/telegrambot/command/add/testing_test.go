package add

import (
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	mocks_storage "gitlab.ozon.dev/veezex/homework/internal/pkg/storage/mocks"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"log"
	"testing"
)

type Fixture struct {
	storage *mocks_storage.MockAppleStorage
	impl    cmdPkg.Command
}

func setUp(t *testing.T) Fixture {
	//t.Parallel()

	stor := mocks_storage.NewMockAppleStorage(gomock.NewController(t))
	return Fixture{
		storage: stor,
		impl:    New(stor),
	}
}

type FakeData struct {
	Color string `faker:"len=25"`
	Price uint64 `faker:"boundary_start=10, boundary_end=10000"`
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
