package main

import "fmt"

func main() {
	// g := 2 == 2
	// h := 1 <= 3
	// i := 4 >= 7
	// j := 3 != 4
	// k := 2 < 9
	// l := 4 > 8

	g := (2 == 2)
	h := (1 <= 3)
	i := (4 >= 7)
	j := (3 != 4)
	k := (2 < 9)
	l := (4 > 8)
	fmt.Printf("%T\t %v\n", g, g)
	fmt.Printf("%T\t %v\n", h, h)
	fmt.Printf("%T\t %v\n", i, i)
	fmt.Printf("%T\t %v\n", j, j)
	fmt.Printf("%T\t %v\n", k, k)
	fmt.Printf("%T\t %v\n", l, l)
}