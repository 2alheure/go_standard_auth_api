package routes

import (
	_ "fmt"
	_ "net/http"

	_ "github.com/gorilla/mux"

	_ "github.com/2alheure/go_standard_auth_api/helpers"
	"github.com/2alheure/go_standard_auth_api/controllers"
)

func (router *MuxRouter) InitAuthRoutes() (*MuxRouter) {
	subrouter := router.MR.PathPrefix("/auth/").Subrouter()
	
	/**
	* @apiDefine paramEmail
	* @apiParam (Mandatory) {string} email An email address
	*/

	/**
	* @apiDefine paramLogin
	* @apiParam (Mandatory) {string} login The login
	*/

	/**
	* @apiDefine paramPassword
	* @apiParam (Mandatory) {string} password The password
	*/
 

	/**
	* @api {get} / AccountInfo
	* @apiDescription Gives information about the currently connected user
	* @apiName AccountInfo
	* @apiGroup Auth
	* @apiUSe TokenNeeded
	*/
	subrouter.HandleFunc("/", controllers.AccountInfo).Methods("GET")
	
	/**
	* @api {post} / Login
	* @apiDescription Logs in the user.
	* @apiName Login
	* @apiGroup Auth
	* @apiUse paramLogin
	* @apiUse paramPassword
	* @apiUse BadParamError
	*/
	subrouter.HandleFunc("/", controllers.Login).Methods("POST")
	
	/**
	* @api {put} / UpdateAccount
	* @apiDescription Modifies account info.
	* @apiName UpdateAccount
	* @apiGroup Auth
	* @apiUSe TokenNeeded
	* @apiParam (Optional) {string} email The new email
	* @apiParam (Optional) {string} login The new login
	* @apiParam (Optional) {string} password The new password
	* @apiUse BadParamError
	*/
	subrouter.HandleFunc("/", controllers.AccountUpdate).Methods("PUT")
	
	/**
	* @api {delete} / DeleteAccount
	* @apiDescription Deletes the account.
	* @apiName DeleteAccount
	* @apiGroup Auth
	* @apiUSe TokenNeeded
	*/
	subrouter.HandleFunc("/", controllers.DeleteAccount).Methods("DELETE")

	/**
	* @api {post} /register Register
	* @apiDescription Registers the account.
	* @apiName Register
	* @apiGroup Auth
	* @apiUse paramLogin
	* @apiUse paramPassword
	* @apiUse paramEmail
	* @apiUse BadParamError
	*/
	subrouter.HandleFunc("/register", controllers.Register).Methods("POST")
	
	/**
	* @api {post} /recover Recover
	* @apiDescription Sends a mail to recover the account.
	* @apiName Recover
	* @apiGroup Auth
	* @apiUse paramEmail
	* @apiUse BadParamError
	*/
	subrouter.HandleFunc("/recover", controllers.Recover).Methods("POST")

	return router
}