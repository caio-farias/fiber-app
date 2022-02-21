package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey", json: "id"`
	Firstname string `gorm:not null, json: "firstName"`
	Lastname string `gorm:not null, json: "lastName"`
	Email string `gorm:"unique", json: "email"`
	Username string `gorm:"unique", json: "username"`
	Password string `gorm:not null, json: "password"`
	CreatedAt time.Time `json: "createdAt"`
  UpdatedAt time.Time `json: "updatedAt"`
  DeletedAt gorm.DeletedAt `gorm:"index", json: "deletedAt"`
}