package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	commanderPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot"
	cmdAddPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/add"
	cmdDelPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/delete"
	cmdGetPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/get"
	cmdListPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command/update"
	routerPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/router"
)

func runTelegramBot(s storage.AppleStorage, apiKey string, rdb *redis.Client, pubSubChanel string) {
	telegrambot, err := commanderPkg.New(apiKey)
	if err != nil {
		logrus.Fatal(err)
	}

	// subscribe to redis
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		pubsub := rdb.Subscribe(ctx, pubSubChanel)
		defer pubsub.Close()

		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println(msg.Channel, msg.Payload)

			queryId := msg.Payload

			value, err := rdb.HGet(ctx, queryId, "value").Result()
			if err != nil {
				logrus.Error(err)
			}

			meta, err := rdb.HGet(ctx, queryId, "meta").Result()
			if err != nil {
				logrus.Error(err)
			}

			var list []applePkg.Apple
			if meta == "list" {
				err := json.Unmarshal([]byte(value), &list)
				if err != nil {
					logrus.Error(err)
					continue
				}
			}

			if meta == "add" {
				var apple applePkg.Apple
				err := json.Unmarshal([]byte(value), &apple)
				if err != nil {
					logrus.Error(err)
					continue
				}

				list = []applePkg.Apple{apple}
			}

			for _, apple := range list {
				_, err = telegrambot.Send(apple.String())
				if err != nil {
					logrus.Error(err)
					continue
				}
			}
		}
	}()

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
