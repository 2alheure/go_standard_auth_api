package controllers

import (
	"net/http"

	"github.com/2alheure/go_standard_auth_api/helpers"
	"github.com/2alheure/go_standard_auth_api/models"
)


func AccountInfo(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.CheckToken(r)

	var msg map[string]interface{}
	if err != nil {
		msg = helpers.TokenErrorMessage(err)
	} else {
		helpers.RewriteToken(w, r)
		msg = models.AccountInfo(token.UserID)
	}

	helpers.Respond(w, msg)
}

func Login(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("login", "password")

	paramError := helpers.CheckParams(r, nil, post)

	var msg map[string]interface{}
	if paramError != nil {
		msg = helpers.BadParamMessage(paramError)
	} else {
		userID, isAuth := models.Login(r.FormValue("login"), helpers.HashPassword(r.FormValue("password")))
		
		if isAuth {
			token, err := helpers.CreateToken(userID)
			if err != nil {
				msg = helpers.TokenErrorMessage(err)
			} else {
				msg = helpers.OKMessage()
				msg["token"] = token
				helpers.WriteToken(w, token)
			}
		} else {
			msg = helpers.Message(false, 403, "Error in authentification.")
		}
	}

	helpers.Respond(w, msg)
}

func AccountUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.CheckToken(r)

	var msg map[string]interface{}
	if err != nil {
		msg = helpers.ErrorMessage(err)
	} else {
		helpers.RewriteToken(w, r)
		msg = models.AccountUpdate(token.UserID, r)
		msg["info"] = models.AccountInfo(token.UserID)
	}

	helpers.Respond(w, msg)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.CheckToken(r)

	var msg map[string]interface{}
	if err != nil {
		msg = helpers.ErrorMessage(err)
	} else {
		msg = models.DeleteAccount(token.UserID)
	}

	helpers.Respond(w, msg)
}

func Register(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("email", "login", "password")

	paramError := helpers.CheckParams(r, nil, post)

	var msg map[string]interface{}
	if paramError != nil {
		msg = helpers.BadParamMessage(paramError)
	} else {
		msg = models.Register(r.FormValue("email"), r.FormValue("login"), r.FormValue("password"))
	}

	helpers.Respond(w, msg)
}

func Recover(w http.ResponseWriter, r *http.Request) {
	post := new(helpers.Params)
	post.AddMandatory("login")

	paramError := helpers.CheckParams(r, nil, post)

	var msg map[string]interface{}
	if paramError != nil {
		msg = helpers.BadParamMessage(paramError)
	} else {
		_, isFound := models.Recover(r.FormValue("login"))

		if isFound {
			var mailSent bool
			// Here, we need to send an email to the address of the user
			// with a link leading to a reset password form
			// or giving a token to do so

			if mailSent {
				msg = helpers.Message(true, 200, "An email has been sent to the email linked to the account. Please follow its instructions in order to reset your password.")
			} else {
				msg = helpers.Message(false, 500, "An unexpected error happened while sending a recovering mail.")
			}
		} else {
			msg = helpers.Message(false, 404, "No account exists with this login.")
		}
	}


	helpers.Respond(w, msg)
}
