package config

import (
	"github.com/jinzhu/configor"
	"telegram-file-server/pkg/log"
)

type environment struct {
	ServerHost          string `env:"SERVER_HOST" yaml:"ServerHost" default:"0.0.0.0"`
	ServerPort          int    `env:"SERVER_PORT" yaml:"ServerPort" default:"8080"`
	ServerBodySizeLimit int64  `env:"SERVER_BODY_SIZE_LIMIT" yaml:"ServerBodySizeLimit" default:"2147483648"` // 2GB
	LogLevel            string `env:"LOG_LEVEL" yaml:"LogLevel" default:"info"`
	Secret              string `env:"SECRET" yaml:"Secret" required:"true"`

	// Telegram
	_                      int    `env:"TELEGRAM_API_ID" required:"true"`
	_                      string `env:"TELEGRAM_API_HASH" required:"true"`
	TelegramBotToken       string `env:"TELEGRAM_BOT_TOKEN" yaml:"TelegramBotToken" required:"true"`
	TelegramChatID         int64  `env:"TELEGRAM_CHAT_ID" yaml:"TelegramChatID" required:"true"`
	TelegramBotApiEndpoint string `env:"TELEGRAM_BOT_API_ENDPOINT" yaml:"TelegramBotApiEndpoint" default:"http://localhost:8081/bot%s/%s"`
}

var Environment environment

func init() {
	log.Info("Server is starting up...")
	log.Panic(configor.Load(&Environment, "config.yml"))
	if len(Environment.Secret) != 32 {
		log.Error("Secret must be 32 characters long")
		panic("Secret must be 32 characters long")
	}
}
