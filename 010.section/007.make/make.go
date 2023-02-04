package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 9
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[10] = 99999
	x = append(x, 99999)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 11) // cap가 12에서 24로 2배로 늘어나게 됨
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := []int{1, 2, 3}
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}
