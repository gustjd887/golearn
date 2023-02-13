package main

import "fmt"

// func main() {
// 	defer fmt.Println(foo())
// 	fmt.Println(bar())
// 	fmt.Println(too())
// }

// func foo() string {
// 	return "foo"
// }

// func bar() string {
// 	return "bar"
// }

// func too() string {
// 	return "too"
// }

func main() {
	defer foo()
	fmt.Println("main func")
}

func foo() {
	defer func() {
		fmt.Println("defer foo func")
	}()
	fmt.Println("foo func")
}
