package main

import (
	"github.com/faryne/chat/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	f := fiber.New(fiber.Config{
		ServerHeader:  "",
		StrictRouting: true,
		CaseSensitive: true,
	})
	routes.InitWebsocketRoutes(f)

	_ = f.Listen(":8080")
}
