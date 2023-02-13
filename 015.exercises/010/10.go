package main

import "fmt"

func main() {
	x := 1
	{
		y := 2
		fmt.Println(y)
	}
	fmt.Println(x)
	// fmt.Println(y)
}
