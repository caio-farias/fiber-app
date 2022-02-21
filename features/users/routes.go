package users

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router){
	usersRoutes := router.Group("/users")
	usersRoutes.Get("", GetUsers)
	usersRoutes.Post("", CreateUser)
	usersRoutes.Get("/:id", GetUser)
	usersRoutes.Patch("/:id", UpdateUser)
	usersRoutes.Delete("/:id", DeleteUser)
}

