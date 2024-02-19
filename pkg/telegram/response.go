package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type SendDocumentResult struct {
	tgbotapi.MessageID
	Document tgbotapi.Document `json:"document"`
}

type GetFileResult struct {
	tgbotapi.File
}
