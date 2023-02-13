package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(i...))
	fmt.Println(bar(i))
}

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
