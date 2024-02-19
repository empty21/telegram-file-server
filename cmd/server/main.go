package main

import (
	"github.com/gofiber/fiber/v2"
	f "telegram-file-server/pkg/fiber"
	"telegram-file-server/pkg/route"
)

func main() {
	// Start the server
	app := fiber.New(f.Config())

	route.RegisterAllRoutes(app)

	f.StartServer(app)
}
