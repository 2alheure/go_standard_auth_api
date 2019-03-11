package controllers

import (
	"net/http"

	"github.com/2alheure/go_standard_auth_api/helpers"
)


func AccountInfo(w http.ResponseWriter, r *http.Request) {}

func Login(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("id")

	getErr, postErr := helpers.CheckParams(r, nil, post)

	helpers.HandleError(getErr)
	helpers.HandleError(postErr)
}

func AccountUpdate(w http.ResponseWriter, r *http.Request) {}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {}

func Register(w http.ResponseWriter, r *http.Request) {}

func Recover(w http.ResponseWriter, r *http.Request) {}
