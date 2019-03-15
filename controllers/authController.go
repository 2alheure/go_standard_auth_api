package controllers

import (
	"net/http"

	"github.com/2alheure/go_standard_auth_api/helpers"
)


func AccountInfo(w http.ResponseWriter, r *http.Request) {}

func Login(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("login", "password")

	paramError := helpers.CheckParams(r, nil, post)

	var msg map[string]interface{}
	if paramError != nil {
		msg = helpers.BadParamMessage(paramError)
	} else {
		msg = helpers.OKMessage()
	}

	helpers.Respond(w, msg)
}

func AccountUpdate(w http.ResponseWriter, r *http.Request) {}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {}

func Register(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("email", "login", "password")
	get := new(helpers.Params)
	get.AddMandatory("test", "login", "password")

	paramError := helpers.CheckParams(r, get, post)

	var msg map[string]interface{}
	if paramError != nil {
		msg = helpers.BadParamMessage(paramError)
	} else {
		msg = helpers.OKMessage()
	}

	helpers.Respond(w, msg)
}

func Recover(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("email")

	paramError := helpers.CheckParams(r, nil, post)

	var msg map[string]interface{}
	if paramError != nil {
		msg = helpers.BadParamMessage(paramError)
	} else {
		msg = helpers.OKMessage()
	}

	helpers.Respond(w, msg)
}
