package main

import (
	"fmt"
	"os"
)

type customErr struct {
	msg string
	err error
}

func (e customErr) Error() string {
	return fmt.Sprintf("%v %v", e.msg, e.err)
}

func foo(e error) {
	fmt.Printf("%v\n%v\n", e.(customErr).msg, e.(customErr).err)
}

func main() {
	// err := fmt.Errorf("custom err msg")
	// fmt.Printf("%T\n", err)
	_, err := os.ReadFile("no_file.txt")
	if err != nil {
		cs := customErr{
			"custom error msg",
			err,
		}
		foo(cs)
	}
}
