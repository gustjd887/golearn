package main

import (
	"errors"
	"fmt"
	"log"
)

var errNorgate error = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", errNorgate)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errNorgate
	}
	return 42, nil
}
