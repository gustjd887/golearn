package main

import "fmt"

var x int
var y float64

func main () {
	x = 4294967295
	y = 42.324543
	z := 22.22

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(z)

	z = 2
	fmt.Println(z)
	fmt.Printf("%T\n",z)
}