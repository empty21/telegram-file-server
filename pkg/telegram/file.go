package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-file-server/pkg/config"
)

func UploadFile(file tgbotapi.RequestFileData) (string, error) {
	document := tgbotapi.NewDocument(config.Environment.TelegramChatID, file)
	message, err := bot.Send(document)
	if err != nil {
		return "", err
	}
	return message.Document.FileID, nil
}

func GetFile(fileId string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{
		FileID: fileId,
	})
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return file.FilePath, nil
}
