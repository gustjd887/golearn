package main

import "fmt"

func main() {
	f1 := returnFunc()
	fmt.Println(f1([]int{1, 2, 3}...))
}

func returnFunc() func(i ...int) int {
	return func(i ...int) int {
		total := 0
		for _, v := range i {
			total += v
		}
		return total
	}
}
