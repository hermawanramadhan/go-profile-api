package app

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = "8080"
)

func StartApp() {

	webServerMsg := fmt.Sprintf("Start development server localhost:%v", port)
	log.Println(webServerMsg)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), route()))

}
