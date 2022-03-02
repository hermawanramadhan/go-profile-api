package app

import (
	"customer-profile/config"

	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func StartApp() {

	webServerMsg := fmt.Sprintf("Start development server localhost:%v", config.Http_port)
	log.Println(webServerMsg)
	handler := cors.AllowAll().Handler(router())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Http_port), handler))

}
