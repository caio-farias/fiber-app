package config

import (
	"fiber-app/features/users"

	"github.com/gofiber/fiber/v2"
)


func SetupRouter(router fiber.Router){
	users.SetupUserRoutes(router)
}