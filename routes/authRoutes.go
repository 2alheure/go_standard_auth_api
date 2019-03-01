package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/2alheure/go_standard_auth_api/helpers"
	"github.com/2alheure/go_standard_auth_api/controllers"
)

func (router http.Handler) InitAuthRoutes() (router http.Handler) {
	subrouter := router.PathPrefix("/auth/").Subrouter()

	/**
	* @api {get} / Auth
	* @apiDescription Gives info about the currently connected account.
	* @apiGroup Auth
	*/
	subrouter.HandleFunc("/", controllers.AccountInfo).Methods("GET")
	
	/**
	* @api {post} / Auth
	* @apiDescription Logs in the user.
	* @apiGroup Auth
	*/
	subrouter.HandleFunc("/", controllers.Login).Methods("POST")
	
	/**
	* @api {put} / Auth
	* @apiDescription Modifies account info.
	* @apiGroup Auth
	*/
	subrouter.HandleFunc("/", controllers.AccountUpdate).Methods("PUT")
	
	/**
	* @api {delete} / Auth
	* @apiDescription Deletes the account.
	* @apiGroup Auth
	*/
	subrouter.HandleFunc("/", controllers.DeleteAccount).Methods("DELETE")

	/**
	* @api {post} / Auth
	* @apiDescription Registers the account.
	* @apiGroup Auth
	*/
	subrouter.HandleFunc("/register", controllers.Register).Methods("POST")
	
	/**
	* @api {post} / Auth
	* @apiDescription Sends a mail to recover the account.
	* @apiGroup Auth
	*/
	subrouter.HandleFunc("/recover", controllers.Recover).Methods("POST")

	return router
}