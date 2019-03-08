package controllers

import (
	"net/http"

	"github.com/2alheure/go_standard_auth_api/helpers"
)


func AccountInfo(w http.ResponseWriter, r *http.Request) {}

func Login(w http.ResponseWriter, r *http.Request) {
	wanted := []string{"id"}
	helpers.CheckParams(r, wanted, nil)
}

func AccountUpdate(w http.ResponseWriter, r *http.Request) {}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {}

func Register(w http.ResponseWriter, r *http.Request) {}

func Recover(w http.ResponseWriter, r *http.Request) {}
