package main

import "fmt"

type phs int
var x phs

func main () {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}