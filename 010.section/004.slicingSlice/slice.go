package main

import "fmt"

func main() {
	x := []int{4,5,6,7,1}
	fmt.Println(len(x))
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) // 1 <= x < 3

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	for i := 0; i < 5; i++ {
		fmt.Println(i, x[i])
	}
}