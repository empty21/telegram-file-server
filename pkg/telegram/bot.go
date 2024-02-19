package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"telegram-file-server/pkg/config"
)

var bot *tgbotapi.BotAPI

func defaultBot() *tgbotapi.BotAPI {
	b := &tgbotapi.BotAPI{
		Token:  config.Environment.TelegramBotToken,
		Debug:  false,
		Buffer: 100,
		Self:   tgbotapi.User{},
		Client: new(http.Client),
	}
	b.SetAPIEndpoint(config.Environment.TelegramBotApiEndpoint)
	return b
}

func init() {
	bot, _ = tgbotapi.NewBotAPIWithAPIEndpoint(config.Environment.TelegramBotToken, config.Environment.TelegramBotApiEndpoint)
	if bot == nil {
		bot = defaultBot()
	}
}
