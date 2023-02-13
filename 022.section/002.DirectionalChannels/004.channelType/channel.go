package main

import "fmt"

func main() {
	c := make(chan int, 2) // 송수신 체널
	cr := make(<-chan int) //수신 전용 체널
	cs := make(chan<- int) //송신 전용 체널

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// // 구체화된 타입(송/수신 전용 체널)은 일반 타입에 할당 불가
	// c = cr
	// c = cs

	// 일반 체널을 구체화된 타입(송/수신 전용 체널)에 할당은 가능
	cr = c
	cs = c

	fmt.Printf("%T\n", (<-chan int)(c))
	fmt.Printf("%T\n", (chan<- int)(c))
}
