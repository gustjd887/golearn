package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// fmt.Println(&c)
	go foo(c) // 밑에 있는 for문이랑 별개로, 계속 체널에 값을 송신하는 역할.

	for v := range c { // range에서 체널 c를 사용하면, 체널이 닫힐 때 까지 체널에서 값을 뺀다.
		fmt.Println(v)
	}
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
