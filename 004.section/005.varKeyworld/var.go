package main

import "fmt"

// x := 42 // 함수 밖에서는 에러 발생
var y = 43 // 함수 밖에서도 에러 발생하지 않음
var z int // 값을 지정하지 않은 int type z 변수 선언
func main() {
	// short declaration
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()
	fmt.Println("int type을 지정한 z 출력 = ",z) // 값을 지정하지 않으면 0이 출력됨(default)
}

func foo() {
	fmt.Println("again:", y)
}