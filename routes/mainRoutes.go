package routes

import (
	_ "fmt"
	_ "net/http"

	"github.com/gorilla/mux"

	_ "github.com/2alheure/go_standard_auth_api/helpers"
)

type MuxRouter struct {
	MR		*mux.Router
}

func InitRouter() *MuxRouter {
	router := &MuxRouter{mux.NewRouter()}

	router.InitAuthRoutes()		// Return the router so chainable

	return router
}