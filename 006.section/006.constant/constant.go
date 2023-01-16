package main

import "fmt"

// const a = 42
// const b = 24.78
// const c = "James Bond"
const (
	a = 42
	b = 24.78
	c = "James Bond"
)


func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
}