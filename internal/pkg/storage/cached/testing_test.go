package cached

import (
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	mocks_cache "gitlab.ozon.dev/veezex/homework/internal/pkg/cache/mocks"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	mocks_storage "gitlab.ozon.dev/veezex/homework/internal/pkg/storage/mocks"
	"log"
	"testing"
)

type Fixture struct {
	storage *mocks_storage.MockAppleStorage
	cacher  *mocks_cache.MockCacher
	apple   apple.Apple
	impl    storage.AppleStorage
}

func setUp(t *testing.T) Fixture {
	//t.Parallel()

	cacher := mocks_cache.NewMockCacher(gomock.NewController(t))
	stor := mocks_storage.NewMockAppleStorage(gomock.NewController(t))
	return Fixture{
		impl:    New(stor, cacher),
		cacher:  cacher,
		storage: stor,
		apple: apple.Apple{
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
	//Key   string  `faker:"len=25"`
	//Value string  `faker:"sentence"`
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
