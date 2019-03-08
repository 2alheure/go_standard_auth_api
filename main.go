package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/2alheure/go_standard_auth_api/routes"
	"github.com/2alheure/go_standard_auth_api/helpers"
)

func main() {
	err := godotenv.Load("my.env")
	helpers.HandleError(err)

	router := routes.InitRouter()
	port := os.Getenv("PORT")
	address := os.Getenv("ADDRESS")

	log.Fatal(http.ListenAndServe(address+":"+port, router.MR))
}