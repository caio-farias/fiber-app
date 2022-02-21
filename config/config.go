package config

import (
	"fiber-app/features/users"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func  SetupDb() *gorm.DB {
	dsn := MountDbURL()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("Cannot connect to database..")
	}

	return db
}

func SetAppConfig(app *fiber.App){
	app.Use(logger.New())
	app.Use(cors.New())
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured loading env variables. Err: %s", err)
	}
	db := SetupDb()

	users.InitMigration(db)
	appRoutes := app.Group(os.Getenv("CONTEXT_PATH"))
	SetupRouter(appRoutes)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

