package middleware

import (
	"github.com/gofiber/fiber/v2"
	"telegram-file-server/pkg/config"
	"telegram-file-server/pkg/model"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check if user is authenticated
	// If not, return 401
	// If yes, continue
	apiKey := c.Get("X-Api-Key")
	if apiKey != config.Environment.Secret {
		return c.Status(401).JSON(model.NewErrorResponse("Unauthorized"))
	}
	return c.Next()
}
