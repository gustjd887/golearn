package main

import "fmt"

type phs int
var x phs
var y int
// xy := "xy"

func main () {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// fmt.Println(xy)
}