package helpers

import (
	"fmt"
	"net/http"
	_ "net/url"
)

// Params will handle request parameters expectations
type Params struct {
	Mandatory		[]string		`json:mandatory,omitempty`
	Optionnal		[]string		`json:optionnal,omitempty`
}

// ParamError will handle all parameters errors
// Miss will be filled with names of missing parameters
// Extra will be filled with names of parameters which weren't asked to be there
type ParamError struct {
	Miss			[]string		`json:miss,omitempty`
	Extra			[]string		`json:extra,omitempty`
}

func (err ParamError) Error() string {
	return "Bad request parameters."
}

// AddMandatory is here to add args to mandatory parameters
func (params *Params) AddMandatory(args ...string) {
	params.Mandatory = append(params.Mandatory, args...)
}

// AddOptionnal is here to add args to optionnal parameters
func (params *Params) AddOptionnal(args ...string) {
	params.Optionnal = append(params.Optionnal, args...)
}

// CheckParams will search through httpRequest.Form to examine given POST parameters
// and through url.value to examine given URL parameters
func CheckParams(r *http.Request, get *Params, post *Params) (getErrors *ParamError, postErrors *ParamError) {
	getErrors = new(ParamError)
	postErrors = new(ParamError)

	if len(get.Mandatory) != 0 || len(get.Optionnal) != 0 {		// Checking GET params
		fmt.Print(r.URL.Query())
	}
	
	
	if len(post.Mandatory) != 0 || len(post.Optionnal) != 0 {	// Checking POST params
		r.ParseForm()
		reqForm := r.Form

		for key, _ := range r.Form {		// Checks whether key is an extra
			_, ok := InArray(get.Mandatory, key); 
			_, ok2 := InArray(get.Optionnal, key); 
			if !ok && !ok2 {
				getErrors.Extra = append(getErrors.Extra, key)
			}
		}

		for _, val := range post.Mandatory {
			if testVal, ok := reqForm[val] ; ok && len(testVal) > 0 {		// Checks whether key is missing
				postErrors.Miss = append(postErrors.Miss, val)
			}
		}
	}

	return
}