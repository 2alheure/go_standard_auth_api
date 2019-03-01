package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/gorilla/mux"
	_ "github.com/joho/godotenv"

	_ "github.com/2alheure/go_standard_auth_api/models"
	_ "github.com/2alheure/go_standard_auth_api/helpers"
	_ "github.com/2alheure/go_standard_auth_api/controllers"
	"github.com/2alheure/go_standard_auth_api/routes"
)

func main() {
	router := routes.InitRouter()

	// Change address + port for config value
	server := &http.Server{
		Handler:      router.MR,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}