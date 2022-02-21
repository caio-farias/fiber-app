package users

import (
	"gorm.io/gorm"
)

var usersRepository *gorm.DB

func InitMigration(db *gorm.DB){
	usersRepository = db
	db.AutoMigrate(&User{})
}

func findAll(users *[]User){
	usersRepository.Find(&users)
}

func findOneById(user *User, id string) {
	usersRepository.First(&user, id)
}

func findOneByEmail(user *User, email string) {
	usersRepository.First(&user, "email=?", email)
}


func findOneByUsername(user *User, username string)  {
	usersRepository.First(&user, "username=?", username)
}

func create(user *User) error {
 exec :=	usersRepository.Create(&user); 
 if (exec.Error == nil){ 
	 return nil 
	}
	return exec.Error
}

func update(user *User) error {
 exec :=	usersRepository.Save(&user); 
 if (exec.Error == nil){ 
	 return nil 
	}
	return exec.Error
}

func deleteBydId(user *User, id string){
	usersRepository.First(&user, id)
}


