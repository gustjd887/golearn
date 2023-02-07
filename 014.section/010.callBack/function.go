package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1,2,3,4,5,6,7,8,9}...))
	fmt.Println(even(sum, []int{1,2,3,4,5,6,7,8,9}...))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}