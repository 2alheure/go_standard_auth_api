package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/2alheure/go_standard_auth_api/models"
	"github.com/2alheure/go_standard_auth_api/helpers"
	"github.com/2alheure/go_standard_auth_api/controllers"
	"github.com/2alheure/go_standard_auth_api/routes"
)

func main() {
	router := routes.InitRouter()

	// Change address + port for config value
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}