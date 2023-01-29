package main

import "fmt"

func main() {
	var x [5]int
	var y [6]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(y)
}