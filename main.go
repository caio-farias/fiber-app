package main

import (
	"fiber-app/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.SetAppConfig(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
