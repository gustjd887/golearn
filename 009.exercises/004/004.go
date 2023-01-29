package main

import "fmt"

func main() {
	i := 1989
	for {
		if i > 2023 {
			break
		}
		fmt.Printf("%d\t, i")
		i++
	}
}