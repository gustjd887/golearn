package main

import (
	"fmt"
)

type customErr struct {
	msg string
	err error
}

func (e customErr) Error() string {
	return fmt.Sprintf(e.msg, e.err)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	err := fmt.Errorf("custom err msg")
	// fmt.Printf("%T\n", err)
	cs := customErr{
		"custom err msg",
		err,
	}

	foo(cs)
}
