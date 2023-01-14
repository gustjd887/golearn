package main

import "fmt"

func main() {
	x :=42 // 변수 x를 선언하면서 값을 지정
	fmt.Println("단축연산자로 선안한 x 출력", x)
    x = 32 // 이미 선언되어 있는 변수이기 때문에 =만 사용
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
}