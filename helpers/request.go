package helpers

import (
	_ "fmt"
	"net/http"
	_ "net/url"
)

// Params will handle request parameters expectations
type Params struct {
	Mandatory		[]string			`json:"mandatory,omitempty"`
	Optionnal		[]string			`json:"optionnal,omitempty"`
}

// ParamError will handle all parameters errors
// GET will give details about query parameters errors
// POST will give details about body parameters errors
type ParamError struct {
	GET				DetailParamError	`json:"url,omitempty"`
	POST			DetailParamError	`json:"body,omitempty"`
}

// Miss will be filled with names of missing parameters
// Extra will be filled with names of parameters which weren't asked to be there
type DetailParamError struct {
	Miss			[]string			`json:"miss,omitempty"`
	Extra			[]string			`json:"extra,omitempty"`
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
func CheckParams(r *http.Request, get *Params, post *Params) (err *ParamError) {
	err = new(ParamError)


	r.ParseForm()
	reqForm := r.PostForm
	reqURL := r.URL.Query()


	if get != nil && (len(get.Mandatory) != 0 || len(get.Optionnal) != 0) {		// Checking GET params
		for key, _ := range reqURL {		// Checks whether key is an extra
			_, ok := InArray(key, get.Mandatory); 
			_, ok2 := InArray(key, get.Optionnal); 
			if !ok && !ok2 {
				err.GET.Extra = append(err.GET.Extra, key)
			}
		}
		
		for _, val := range get.Mandatory {
			if testVal, ok := reqURL[val]; !ok || len(testVal) == 0 {		// Checks whether key is missing
				err.GET.Miss = append(err.GET.Miss, val)
			}
		}
		
	} else {
		for key, _ := range reqURL {
			err.GET.Extra = append(err.GET.Extra, key)
		}
	}

	
	if post != nil && (len(post.Mandatory) != 0 || len(post.Optionnal) != 0) {	// Checking POST params
		for key, _ := range reqForm {		// Checks whether key is an extra
			_, ok := InArray(key, post.Mandatory); 
			_, ok2 := InArray(key, post.Optionnal); 
			if !ok && !ok2 {
				err.POST.Extra = append(err.POST.Extra, key)
			}
		}
		
		for _, val := range post.Mandatory {
			if testVal, ok := reqForm[val]; !ok || len(testVal) == 0 {		// Checks whether key is missing
				err.POST.Miss = append(err.POST.Miss, val)
			}
		}

	} else {
		for key, _ := range reqForm {
			err.POST.Extra = append(err.POST.Extra, key)
		}
	}

	
	if len(err.GET.Miss) == 0 && len(err.GET.Extra) == 0 && len(err.POST.Miss) == 0 && len(err.POST.Extra) == 0 {
		return nil
	}

	return
}