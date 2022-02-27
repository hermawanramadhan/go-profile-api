package main

import (
	"customer-profile/app"
	"customer-profile/db"
)

func main() {
	db.Connect()
	app.StartApp()
}
