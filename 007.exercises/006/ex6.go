package main

import "fmt"

const (
	_ = iota
	a = 2023 + iota
	b = 2023 + iota
	c = 2023 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}