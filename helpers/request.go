package helpers

import (
	"net/http"
	"fmt"
)

type Params struct {
	Wanted			map[string]string		`json:wanted,omitempty`
	Optionnal		map[string]string		`json:optionnal,omitempty`
	Miss			map[string]string		`json:miss,omitempty`
	Extra			map[string]string		`json:extra,omitempty`
}

func CheckParams(r *http.Request, wanted []string, optionnal []string) (*Params, error) {
	r.ParseForm()

	for key, val := range r.Form {
		fmt.Printf("%s = %s", key, val)
	}

	return nil, nil
}