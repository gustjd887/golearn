package main

import "fmt"

const (
	_ = iota
	// kb = 1024
	// mb = kb * 1024
	// gb = mb * 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
	fmt.Println(kb)
}