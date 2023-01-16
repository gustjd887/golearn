package main

import "fmt"

func main() {
	x := 12
	fmt.Printf("binary: %b\n", x)
	fmt.Printf("demical: %d\n", x)
	fmt.Printf("hexa demical: %#x\n", x)

	y := x << 1
	fmt.Printf("binary: %b\n", y)
	fmt.Printf("demical: %d\n", y)
	fmt.Printf("hexa demical: %#x\n", y)
}