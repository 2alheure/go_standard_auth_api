package models

import (
	"net/http"
	"time"
	"fmt"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/2alheure/go_standard_auth_api/helpers"
)

type User struct {
	// gorm.Model
	ID				int						`gorm:"primary_key",json:"id,omitempty"`
	Pseudo			string					`json:"pseudo,omitempty"`
	Login			string					`json:"login,omitempty"`
	Password		string					`json:"-"`
	Email			string					`json:"email,omitempty"`
	CreatedAt		*time.Time				`json:"signUpDate,omitempty"`
	// Token			*helpers.TokenContent	`json:"token,omitempty"`
}

func AccountInfo(userID int) (map[string]interface{}) {
	user := User{}
	if errors := DB.First(&user).GetErrors() ; len(errors) > 0 {
		fmt.Println(errors)
		return map[string]interface{} {"user":nil}
	} else {
		return map[string]interface{} {"user":user}
	}
}

func Login(login string, password string) (int, bool) {
	user := User{}
	if errors := DB.Where("login=? AND password=?", login, password).First(&user).GetErrors() ; len(errors) > 0 {
		fmt.Println(errors)
		return 0, false
	} else {
		return user.ID, true
	}
}

func AccountUpdate(userID int, r *http.Request) (map[string]interface{}) {
	return map[string]interface{} {}
}

func DeleteAccount(userID int) (bool) {
	user := User{ID: userID}
	if errors := DB.Delete(&user).GetErrors() ; len(errors) > 0 {
		fmt.Println(errors)
		return false
	} else {
		return true
	}
}

func Register(email string, login string, password string) (map[string]interface{}) {
	user := User{Email: email, Login: login, Password: password}

	if errors := DB.Create(&user).GetErrors() ; len(errors) > 0 {
		fmt.Println(errors)
		return helpers.Message(false, 500, "An unexpected error happened while registering.")
	} else {
		return helpers.Message(true, 201, "Registration complete.")
	}
}

func Recover(login string) (User, bool) {
	user := User{Login: login}
	if errors := DB.Where("login=?", login).First(&user).GetErrors() ; len(errors) > 0 {
		fmt.Println(errors)
		return user, false
	} else {
		return user, true
	}
}
