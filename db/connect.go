package db

import (
	"customer-profile/entities"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func Connect() {

	dsn := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", username, password, dbname)
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if Err != nil {
		log.Println("Database Connection failed", Err)
	} else {
		log.Println(fmt.Sprintf("Database connection success. connect to database: %v", dbname))
	}

	DBMigrate()
}

func DBMigrate() {
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.RiskProfile{})
}
