package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	// go bar(c) //
	bar(c) // 보통 체널을 받아서 출력하는 함수에는 go를 붙이면 안될듯...

	fmt.Println("exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
