package models

import (
	"customer-profile/db"
	"customer-profile/entities"
	"customer-profile/utils"

	"gorm.io/gorm"
)

func CreateUser(user *entities.User) *gorm.DB {
	result := db.DB.Create(&user)

	return result
}

func GetUsers(limit int, offset int) []entities.User {

	users := []entities.User{}
	db.DB.
		Select("id, name, age").
		Limit(limit).
		Offset(offset).
		Find(&users)

	return users
}

func GetUser(userId string) (entities.User, *gorm.DB) {
	var user entities.User
	result := db.DB.Select("users.id, users.name, users.age, RiskProfile.*").Joins("RiskProfile").First(&user, userId)

	return user, result
}

func UserPasswordCheck(name, password string) (entities.User, bool) {
	var user entities.User
	result := db.DB.Where("name = ?", name).First(&user)
	if result.RowsAffected == 0 {
		return user, false
	}
	status := utils.CheckPasswordHash(password, *user.Password)

	//hide password field
	user.Password = nil

	return user, status
}
