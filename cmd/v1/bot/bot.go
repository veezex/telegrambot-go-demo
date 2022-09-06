package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	commanderPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot"
	cmdAddPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/add"
	cmdDelPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/delete"
	cmdGetPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/get"
	cmdListPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/update"
	routerPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/router"
)

func runTelegramBot(s storage.AppleStorage, apiKey string) {
	telegrambot, err := commanderPkg.New(apiKey)
	if err != nil {
		logrus.Fatal(err)
	}

	// make a router
	r := routerPkg.New()
	r.RegisterCommand(cmdAddPkg.New(s))
	r.RegisterCommand(cmdDelPkg.New(s))
	r.RegisterCommand(cmdGetPkg.New(s))
	r.RegisterCommand(cmdListPkg.New(s))
	r.RegisterCommand(cmdUpdatePkg.New(s))
	r.RegisterCommand(routerPkg.NewHelpCommand(r))

	telegrambot.Run(r)
}
