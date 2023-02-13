package main

import "fmt"

func main() {
	c := make(chan int, 1) // buffer = 1

	c <- 42 // buffer = 1 로 설정함으로써 체널이 막히지 않음.(머무를 공간이 있다고 표현함.)

	fmt.Println(<-c)
}
