package helpers

type StdErr struct {
	Message		string		`json:"message,omitempty"`
	HTTPCode	int			`json:"code,omitempty"`
}

func (err StdErr) Error() string {
	return err.Message
}

func ErrorMessage(err *StdErr) (map[string]interface{}) {
	return Message(false, err.HTTPCode, err.Error())
}

func BadParamMessage(err *ParamError) (map[string]interface{}) {
	msg := Message(false, 400, err.Error())
	msg["errors"] = err
	return msg
}

func TokenErrorMessage(err error) (map[string]interface{}) {
	return Message(false, 409, err.Error())
}

func LogError(err error) {}