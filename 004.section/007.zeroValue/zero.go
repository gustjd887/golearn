package main

import "fmt"

var y string
var z int

func main() {
	fmt.Println("zero value y 출력: ", y) // 빈 string 출력
	fmt.Printf("%T\n", y)

	y = "값 변경"
	fmt.Println("zero value y 출력: ", y) // 변경된 string y 값 출력
	fmt.Printf("%T\n", y)

	fmt.Println("zero value z 출력: ", z) // 변경된 string y 값 출력
	fmt.Printf("%T\n", z)
}