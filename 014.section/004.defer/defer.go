package main

import "fmt"

func main() {
	defer foo() //func main 함수 끝까지 기다렸다가 실행됨
	bar()
	too()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func too() {
	fmt.Println("too")
}