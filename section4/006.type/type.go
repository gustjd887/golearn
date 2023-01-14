package main

import "fmt"

var y = 42
var z = "string?"
var a = `use back quote and "double quote"`
var aa = `test
ttt
multi line`
func main () {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // y의 타입 출력
	fmt.Printf("%T\n", z) // z의 타입 출력
	fmt.Println(a)
	fmt.Println(aa)


	// z = 43 // var z 로 선언한 이후에는 z = ?? 이런식으로 값 변경 가능.
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}