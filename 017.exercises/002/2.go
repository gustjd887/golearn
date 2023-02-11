package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "ms"
	(*p).last = "l"
}

func main() {
	p1 := person{
		first: "hs",
		last:  "p",
	}
	fmt.Println(p1.last, p1.first)
	changeMe(&p1)
	fmt.Println(p1.last, p1.first)
}
