package main

import "github.com/gofiber/fiber/v3"

func main() {
	f := fiber.New(fiber.Config{
		ServerHeader:  "",
		StrictRouting: true,
		CaseSensitive: true,
	})
	_ = f.Listen(":8080")
}
