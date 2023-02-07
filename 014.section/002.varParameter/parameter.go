package main

import "fmt"

func main() {
	s := foo(1,2,3,4)

	fmt.Println("total is", s)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding,", v, "to the total which is now", sum)
	}
	fmt.Println("The total is,", sum)

	return sum
}