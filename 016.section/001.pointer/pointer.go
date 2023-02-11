package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Println(*&a)

	// var b int = &a
	b := &a
	fmt.Println(*b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &b)

	*b = 43
	fmt.Println(a)
}
