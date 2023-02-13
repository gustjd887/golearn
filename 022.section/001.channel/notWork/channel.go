package main

import "fmt"

func main() {
	c := make(chan int)

	//체널 수신자가 없음
	c <- 42 // 데이터 입력시 체널이 막힘
	// 아마 main 함수가 끝나야 c를 사용할 수 있도록 풀릴 듯?

	fmt.Println(<-c) // 체널이 막혀있기 때문에 불러오지 못함
}
