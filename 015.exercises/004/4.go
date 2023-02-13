package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("my name is", p.last, p.first, ", my age is", p.age)
}

func main() {
	p1 := person{
		first: "hs",
		last:  "p",
		age:   23,
	}

	p1.speak()
}
