package main

import "fmt"

func main() {
	// xi := []int{1,2,3,4,5,6,7}
	// x := sum(xi...)
	x := sum("James")
	fmt.Println("total is", x)
}

func sum(s string, x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}