package main

import "fmt"

func main() {
	c := make(chan int)

	//체널 수신자가 없음. 값 송신 수신은 같은 타이밍에 해야함.
	// main 함수는 흐름이 있기 때문에 block상태로 유지됨
	c <- 42 // 데이터 입력시 체널이 막힘, 수신자가 없어서.

	fmt.Println(<-c) // 순서대로 실행되기 때문에 불러오지 못함
}
