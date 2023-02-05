package main

import "fmt"

func main() {
	a1 := struct {
		name string
		test string
	}{
		name: "phs",
		test: "test",
	}
	fmt.Println(a1.name, a1.test)
}
