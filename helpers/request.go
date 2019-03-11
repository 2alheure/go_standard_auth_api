package helpers

import (
	"net/http"
	"fmt"
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

// Mandatory is here to add args to mandatory parameters
func (params *Params) Mandatory(args ...string) {
	params.Mandatory = append(params.Mandatory, args)
}

// Optionnal is here to add args to optionnal parameters
func (params *Params) Optionnal(args ...string) {
	params.Optionnal = append(params.Optionnal, args)
}

// CheckParams will search through httpRequest.Form to examine given POST parameters
// and through url.value to examine given URL parameters
func CheckParams(r *http.Request, get Params, post Params) (getErrors ParamError, postErrors ParamError) {
	getErrors := ParamError{}
	postErrors := ParamError{}

	if len(get.Mandatory) != 0 || len(get.Optionnal) != 0 {
		query, err := url.ParseQuery(r.QueryString)		// Needed function
	}
	
	
	if len(post.Mandatory) != 0 || len(post.Optionnal) != 0 {
		r.ParseForm()

		for key, val := range r.Form {		// Checks whether key is an extra
			// Need : isset function
			if !isset(get.Mandatory, key) {
				getErrors.Extra = append(getErrors.Extra, key)
			}
		}

		for key, val = range post.Mandatory {
			if testVal := r.Form(key) ; testVal == "" {		// Checks whether key is missing
				postErrors.Miss = append(postErrors.Miss, key)
			}
		}
	}

	return
}