package main

import (
	"fiber-app/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.SetAppConfig(app)
}
