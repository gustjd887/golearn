package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")	
	}
	f()

	x := func(x int) {
		fmt.Println("the year big brother started watching", x)	
	}
	x(1984)
}