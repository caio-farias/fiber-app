package users

import (
	"fiber-app/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	bodyParserError := c.BodyParser(user)

	if bodyParserError != nil{
		c.Status(fiber.StatusBadRequest).SendString(bodyParserError.Error())
		return bodyParserError
	}

	var sameFieldUser User
	findOneByEmail(&sameFieldUser, user.Email)
	if(sameFieldUser.Email == user.Email){
		return c.Status(fiber.StatusBadRequest).SendString(utils.EmailInUse)
	}

	findOneByUsername(&sameFieldUser, user.Username)
	if(sameFieldUser.Username == user.Username){
		return c.Status(fiber.StatusBadRequest).SendString(utils.UsernameInUse)
	}

	repoError := create(user)
	if(repoError != nil){
		return c.Status(fiber.StatusBadRequest).SendString(repoError.Error())
	} 

	return c.Status(fiber.StatusOK).JSON(user)
}


func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	findOneById(&user, id)
	
	return c.JSON(user)
}

 func GetUsers(c *fiber.Ctx) error {
	var users []User
	findAll(&users)
	
	return c.JSON(users)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	
	var user User
	findOneById(&user, id)
	if(user.Email == ""){
		return c.Status(fiber.StatusBadRequest).SendString(utils.NotFoundMessage)	
	}

	if err := c.BodyParser(&user); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	update(&user)
	return c.JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	findOneById(&user, id)
	if(user.Email == ""){
		return c.Status(fiber.StatusBadRequest).SendString(utils.NotFoundMessage)	
	}

	deleteBydId(&user, id)
	return c.SendStatus(fiber.StatusOK)
}