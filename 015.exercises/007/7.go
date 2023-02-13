package main

import "fmt"

var x int
var g func(i ...int) int

func main() {
	s1 := sum
	fmt.Printf("%T\n", s1)
	fmt.Println(s1([]int{1, 2, 3, 4, 2, 2, 3, 3, 54}...))

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", g)
	g = sum
	fmt.Println(g(1, 2, 3, 4))

}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
