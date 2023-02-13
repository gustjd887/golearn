package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok) // 열려있는 체널의 값을 가져오면 v는 42, ok는 true를 출력한다.

	v, ok = <-c
	fmt.Println(v, ok) // 체널이 닫혀있으면 v는 0, ok는 false를 출력한다.
}
