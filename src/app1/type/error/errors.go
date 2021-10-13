package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

var errDivisionByZero = errors.New("division by zero")

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func run2() error {
	return errDivisionByZero
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	if err := run2(); err != nil {
		fmt.Println(err)
	}
}
