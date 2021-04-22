package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func errorFunc() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	ctxFields := logrus.Fields{
		"app":         "go-errors",
		"environment": "dev",
	}

	err := errors.New("Error!")
	if err != nil {
		logrus.WithFields(ctxFields).WithError(err).Error("ErrMSg")
		return err
	}
	return nil
}

func main() {
	err := errorFunc()
	logrus.Error("Error thrown: ", fmt.Sprintf("%v", err))
}
