package main

import (
	"fmt"
)

func main() {
	i := [5]int{1, 2, 3, 4, 5}

	fmt.Println(i[2:4])

	for i, v := range i {
		fmt.Printf("%d %d\t", i, v)
	}
	fmt.Println("")
	fmt.Printf("%T\n", i)

	test := 42
	fmt.Printf("%T\n", test)
}
