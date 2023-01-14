package main

import "fmt"

var y = 42

func main() {
	fmt.Printf("asdfsadf")
	fmt.Printf("%T\n", y) // 타입 출력
	fmt.Printf("%b\n", y) // 2진수
	fmt.Printf("%x\n", y) // 16진수
	fmt.Printf("%#x\n", y) // 0X + 16진수

	fmt.Printf("%#x\t%b\t%x\n", y, y, y) // 0X + 16진수

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // 0X + 16진수
	fmt.Println(s)
	fmt.Printf("%v", y)
}