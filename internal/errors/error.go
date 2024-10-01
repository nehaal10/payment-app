package errors

import (
	"encoding/json"
	"fmt"
)

type Errors struct {
	Operation string
	ErrorCode string
	Message   string
	Error     string
}

func New(op string, errorCode string, msg string, err error) error {
	var errors []*Errors
	if err == nil {
		err := fmt.Errorf("")
		newError := formatError(op, errorCode, msg, err, errors)
		return newError
	}

	newError := formatError(op, errorCode, msg, err, errors)
	return newError
}

func formatError(op string, errorCode string, msg string, prevErr error, errors []*Errors) error {
	errSting := prevErr.Error()
	errors = append(errors, &Errors{
		Operation: op,
		ErrorCode: errorCode,
		Message:   msg,
		Error:     errSting,
	})

	errorByteData, err := json.Marshal(errors)
	if err != nil {
		panic("failed to create an error")
	}

	formatedError := fmt.Errorf(string(errorByteData))
	return formatedError
}
