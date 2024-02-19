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
	"telegram-file-server/pkg/repository"
	"telegram-file-server/pkg/telegram"
)

var fileRepository repository.FileRepository

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
	err = fileRepository.Save(file)

	if err != nil {
		log.Error("Failed to save file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Failed to save file"))
	}

	return c.Status(200).JSON(model.NewResponse("File uploaded successfully", file))
}

func getFile(c *fiber.Ctx) error {
	fileUuid := c.Params("id")
	if fileUuid == "" {
		return c.Status(400).JSON(model.NewErrorResponse("File id is required"))
	}

	file, err := fileRepository.FindById(fileUuid)
	if err != nil {
		log.Error("Failed to get file: %v", err.Error())
		return c.Status(404).JSON(model.NewErrorResponse("File not found"))
	}

	filePath, err := cache.GetFile(telegram.GetFile)(file.FileId)
	if err != nil {
		log.Error("Failed to get file: %v", err.Error())
		return c.Status(500).JSON(model.NewErrorResponse("Failed to get file"))
	}

	c.Set("Content-Disposition", "attachment; filename="+file.Name)
	return c.Status(200).SendFile(filePath)
}

func deleteFile(c *fiber.Ctx) error {
	fileUuid := c.Params("id")
	if fileUuid == "" {
		return c.Status(400).JSON(model.NewErrorResponse("File id is required"))
	}
	err := fileRepository.Delete(fileUuid)
	if err != nil {
		return c.Status(500).JSON(model.NewErrorResponse("Failed to delete file"))
	}
	return c.Status(200).JSON(model.NewResponse("File deleted successfully", nil))
}

func RegisterFileRoutes(app *fiber.App) {
	app.Post("/upload", uploadFile, middleware.AuthMiddleware)
	app.Delete("/:id", deleteFile, middleware.AuthMiddleware)
	app.Get("/:id", getFile)
	fileRepository = repository.NewFileRepository(config.Database)
}
