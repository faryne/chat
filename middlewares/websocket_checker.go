package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/saveblush/gofiber3-contrib/websocket"
)

func WebsocketCheckerMiddleware(c fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
