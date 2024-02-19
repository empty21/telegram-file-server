package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"telegram-file-server/pkg/config"
)

func Config() fiber.Config {
	return fiber.Config{
		BodyLimit:         int(config.Environment.ServerBodySizeLimit),
		StreamRequestBody: false,
	}
}

func StartServer(app *fiber.App) {
	app.Config()
	panic(app.Listen(fmt.Sprintf("%s:%d", config.Environment.ServerHost, config.Environment.ServerPort)))
}
