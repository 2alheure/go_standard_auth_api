package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/2alheure/go_standard_auth_api/helpers"
)

func InitRouter() http.Handler {
	router := mux.NewRouter()

	router.InitAuthRoutes()		// Return the router so chainable

	return router
}