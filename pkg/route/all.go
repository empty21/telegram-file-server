package route

import "github.com/gofiber/fiber/v2"

func RegisterAllRoutes(app *fiber.App) {
	RegisterFileRoutes(app)
}
