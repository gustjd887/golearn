package main

import "fmt"

var a int
type hotdog int
var b hotdog

func main () {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// a = b // 안됨, b의 타입이 hotdog라서..
	// 정적 프로그래밍에서 type을 동적으로 변환할 수 없음..

	
}