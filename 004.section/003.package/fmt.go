package main

import (
	"fmt"
)

func main() {
	// n, err := fmt.Println("스트링", 42, true)
	// fmt.Println(n)
	// fmt.Println(err)

	// n, _ := fmt.Println("스트링", 42, true) // err을 버림
	// fmt.Println(n)


	n, e := fmt.Println("스트링", 42, true) // e 변수 선언 후 사용하지 않는 경우
	fmt.Println(n)
    // 결과적으로 에러남
}
