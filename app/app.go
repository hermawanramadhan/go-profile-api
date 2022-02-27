package app

import (
	"customer-profile/config"

	"fmt"
	"log"
	"net/http"
)

func StartApp() {

	webServerMsg := fmt.Sprintf("Start development server localhost:%v", config.Http_port)
	log.Println(webServerMsg)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Http_port), route()))

}
