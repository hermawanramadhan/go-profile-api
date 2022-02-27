package models

import (
	"customer-profile/db"
	"customer-profile/entities"

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
