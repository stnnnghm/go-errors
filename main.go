package main

import (
	"fmt"
)

// CustomError will be type error
type CustomError struct{}

// Error satisfies error interface
func (custErr *CustomError) Error() string {
	return "Error!"
}

func main() {
	// create error
	err := &CustomError{}

	fmt.Println(err)

	fmt.Printf("Type of err is %T \n", err)
	fmt.Printf("Value of err is %#v \n", err)
}
