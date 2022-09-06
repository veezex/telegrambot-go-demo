package telegrambot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/router"
)

type Commander interface {
	Run(router router.CommandExecutor) error
	Send(msg string) (tgbotapi.Message, error)
}

type commander struct {
	bot        *tgbotapi.BotAPI
	lastChatId int64
}

func New(apiKey string) (Commander, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "Init tgbot")
	}

	bot.Debug = true
	logrus.Printf("Authorized on account %s", bot.Self.UserName)

	return &commander{
		bot: bot,
	}, nil
}

func (c *commander) Send(msg string) (tgbotapi.Message, error) {
	return c.bot.Send(tgbotapi.NewMessage(c.lastChatId, msg))
}

func (c *commander) Run(router router.CommandExecutor) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		c.lastChatId = update.Message.Chat.ID

		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			res, err := router.HandleCommand(cmd, update.Message.CommandArguments())
			if err != nil {
				msg.Text = err.Error()
			}
			if res == "" {
				continue
			}
			msg.Text = res
		} else {
			logrus.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf("you sent <%v>", update.Message.Text)
		}
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
