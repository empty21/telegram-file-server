package route

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"os"
	"path"
	"telegram-file-server/pkg/cache"
	"telegram-file-server/pkg/config"
	"telegram-file-server/pkg/log"
	"telegram-file-server/pkg/middleware"
	"telegram-file-server/pkg/model"
	"telegram-file-server/pkg/telegram"
	"telegram-file-server/pkg/util"
)

var fileUtil util.FileUtil

func uploadFile(c *fiber.Ctx) error {
	requestFile, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(model.NewErrorResponse("File is required"))
	}
	tempDir, err := os.MkdirTemp("", "tmp-")
	tempPath := path.Join(tempDir, requestFile.Filename)
	err = c.SaveFile(requestFile, tempPath)

	if err != nil {
		log.Error("Failed to save file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Failed to save file"))
	}
	defer func() {
		_ = os.RemoveAll(tempPath)
	}()
	// Upload file to telegram
	fileId, err := telegram.UploadFile(tgbotapi.FilePath(tempPath))
	if err != nil {
		log.Error("Failed to upload file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Failed to upload file"))
	}
	file := model.NewFile(fileId, requestFile.Filename, requestFile.Header.Get("Content-Type"))

	err = fileUtil.EncryptFile(file)
	if err != nil {
		log.Error("Failed to encrypt file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Internal server error"))

	}

	return c.Status(200).JSON(model.NewResponse("File uploaded successfully", file))
}

func getFile(c *fiber.Ctx) error {
	fileId := c.Params("id")
	if fileId == "" {
		return c.Status(400).JSON(model.NewErrorResponse("File id is required"))
	}
	file := model.NewEncryptedFile(fileId)
	err := fileUtil.DecryptFile(file)

	if err != nil {
		log.Error("Failed to decrypt file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Failed to decrypt file"))
	}

	log.Info("File: %v", file.Name)

	filePath, err := cache.GetFile(telegram.GetFile)(file.FileId)
	if err != nil {
		log.Error("Failed to get file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Failed to get file"))
	}

	c.Set("Content-Disposition", "attachment; filename="+file.Name)
	return c.Status(200).SendFile(filePath)
}

func RegisterFileRoutes(app *fiber.App) {
	app.Post("/upload", middleware.AuthMiddleware, uploadFile)
	app.Get("/:id", getFile)
	fileUtil = util.NewFileUtil(config.Environment.Secret)
	if fileUtil == nil {
		log.Error("Failed to create file util")
	}
}
