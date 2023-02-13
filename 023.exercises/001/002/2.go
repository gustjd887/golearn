package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
