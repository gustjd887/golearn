package main

import "fmt"

func main() {

	bt := []byte("ABC")
	fmt.Printf("%T\n", bt) // []byte로 선언하면, type은 []uint8로 찍힘.
}
