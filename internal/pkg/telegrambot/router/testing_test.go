package router

import (
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	mocks_command "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/mocks"
	"log"
	"testing"
)

type Fixture struct {
	command *mocks_command.MockCommand
	router  Router
}

func setUp(t *testing.T) Fixture {
	//t.Parallel()

	return Fixture{
		command: mocks_command.NewMockCommand(gomock.NewController(t)),
		router:  New(),
	}
}

type FakeData struct {
	CommandName   string `faker:"len=25"`
	CommandDesc   string `faker:"sentence"`
	CommandArgs   string `faker:"sentence"`
	CommandResult string `faker:"sentence"`
}

func fake() FakeData {
	fd := FakeData{}
	if err := faker.FakeData(&fd); err != nil {
		log.Fatal(err)
	}
	return fd
}
