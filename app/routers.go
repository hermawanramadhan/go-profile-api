package app

import (
	"customer-profile/controllers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Root)
	router.HandleFunc("/user", controllers.CreateUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("OPTIONS", "GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("OPTIONS", "GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("OPTIONS", "GET")

	router.HandleFunc("/login", controllers.Login).Methods("OPTIONS", "POST")

	return router
}
