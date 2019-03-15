package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/2alheure/go_standard_auth_api/routes"
)

func main() {
	err := godotenv.Load("my.env")
	if err != nil {
		log.Fatal("Unable to access .env variables.")
	}

	router := routes.InitRouter()
	port := os.Getenv("PORT")
	address := os.Getenv("ADDRESS")

	log.Fatal(http.ListenAndServe(address+":"+port, router.MR))
}