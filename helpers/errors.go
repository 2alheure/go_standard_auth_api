package helpers

type StdErr struct {
	Message		string		`json:"message,omitempty"`
	HTTPCode	int			`json:"code,omitempty"`
}

func (err StdErr) Error() string {
	return err.Message
}

func BadParamMessage(err *ParamError) (map[string]interface{}) {
	msg :=  Message(false, 400, err.Error())
	msg["errors"] = err
	return msg
}

func LogError(err error) {}