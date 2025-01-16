package routes

import (
	"github.com/faryne/chat/middlewares"
	"github.com/gofiber/fiber/v3"
)

func InitWebsocketRoutes(f *fiber.App) {
	f.Use("/ws", middlewares.WebsocketCheckerMiddleware)

}
