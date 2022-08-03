package main

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) error {
	if s == "" {
		return errors.New("字符串不能为空")
	}

	return nil
}

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("e.What: %v e.When %v\n", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	// check("hello")
	err := check("")
	fmt.Printf("err.Error(): %v\n", err.Error())

	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
