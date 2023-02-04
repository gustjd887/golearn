package main

import "fmt"

func main() {
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(i[:5])
	fmt.Println(i[5:])
	fmt.Println(i[2:7])
	fmt.Println(i[1:6])
}
