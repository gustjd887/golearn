package main

import "fmt"

func main() {
	x := 42
	y := 43
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
