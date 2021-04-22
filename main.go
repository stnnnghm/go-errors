package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HttpError struct {
	status     int
	httpMethod string
}

// HttpError implements Error method
func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Request %v failed. Server returned %v status", httpError.httpMethod, httpError.status)
}

func GetServerResponse() (string, error) {
	// return mocked error
	return "", &HttpError{401, "GET"}
}

func throwAnotherError() error {
	return errors.New("threw another error")
}

func throwError() error {
	err := throwAnotherError()
	if err != nil {
		return errors.Wrap(err, "Error occurred while processing function throwError")
	}
	return nil
}

func main() {
	err := throwError()
	if err != nil {
		logrus.Error("Error occurred:", fmt.Sprintf("%+v", err))
	} else {
		// Handle server response
		response, err := GetServerResponse()
		if err != nil {
			fmt.Println(err)
			// errval returns a pointer to the instance of err as HttpError
			errval := err.(*HttpError)
			fmt.Printf("erroring method %v", errval.httpMethod)
		} else {
			fmt.Printf("server response %v", response)
		}
	}
}
