package main

import "fmt"

func main() {
	i := 2
	switch {
	case i == 1:
		fmt.Println("i is 1")
	case i == 2:
		fmt.Println("i is 2")
	}
}
