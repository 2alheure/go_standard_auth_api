package helpers

import (
	"log"
)

type StdErr struct {
	Message		string		`json:message,omitempty`
	HTTPCode	int			`json:code,omitempty`
}

func (err StdErr) Error() string {
	return err.Message
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}