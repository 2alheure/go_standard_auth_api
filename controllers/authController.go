package controllers

import (
	"net/http"

	"github.com/2alheure/go_standard_auth_api/models"
	"github.com/2alheure/go_standard_auth_api/helpers"
)


func AccountInfo(w http.ResponseWriter, r *http.Request) {
	token, err := models.CheckToken(r)

	var msg map[string]interface{}
	if err != nil {
		msg = helpers.TokenErrorMessage(err)
	} else {
		helpers.RewriteToken(w, r)
		msg = helpers.OKMessage()
		msg["info"] = models.AccountInfo(token.UserID)["user"]
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
		userID, isAuth := models.Login(r.FormValue("login"), r.FormValue("password"))
		
		if isAuth {
			token, err := models.CreateToken(userID)
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
	token, err := models.CheckToken(r)

	var msg map[string]interface{}
	if err != nil {
		msg = helpers.ErrorMessage(err)
	} else {
		post := new(helpers.Params)
		post.AddOptional("login", "password", "email")
	
		paramError := helpers.CheckParams(r, nil, post)
	
		if paramError != nil {
			msg = helpers.BadParamMessage(paramError)
		} else {
			helpers.RewriteToken(w, r)
			isUpdated := models.AccountUpdate(token.UserID, r)

			if isUpdated {
				msg = helpers.Message(true, 200, "Update succeed.")
				msg["info"] = models.AccountInfo(token.UserID)["user"]
			} else {
				msg = helpers.Message(false, 500, "An unexpected error happened while updating. Changes may not have been done.")
			}
		}
	}

	helpers.Respond(w, msg)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	token, err := models.CheckToken(r)

	var msg map[string]interface{}
	if err != nil {
		msg = helpers.ErrorMessage(err)
	} else {
		accountIsDeleted := models.DeleteAccount(token.UserID)

		if accountIsDeleted {
			msg = helpers.Message(true, 200, "Account deleted successfully.")
		} else {
			msg = helpers.Message(false, 500, "An unexpected error occured while deleting your account. It may have not been deleted.")
		}
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
		msg = models.Register(r.FormValue("email"), r.FormValue("login"), helpers.HashPassword(r.FormValue("password")))
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
