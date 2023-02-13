package main

import "fmt"

func main() {
	fmt.Println(even(sum, []int{1, 2, 3, 4, 4, 2, 3, 3, 21, 2}...))
}

func even(f func(i ...int) int, xi ...int) int {
	vi := []int{}
	for _, v := range xi {
		if v%2 == 0 {
			vi = append(vi, v)
		}
	}
	return f(vi...)
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
