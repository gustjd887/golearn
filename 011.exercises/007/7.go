package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	fmt.Println(x)
	fmt.Println(x[0][0])

	for i, v := range x {
		for j, v := range v {
			fmt.Printf("index %d %d is %v\n", i, j, v)
		}
	}
}
