package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- sepak secretAgent")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- sepak person")
}

type human interface {
	speak()
}

func bar(h human){
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	// sa3 := person{
	// 	first: "hs",
	// 	last: "p",
	// }
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	// sa3.speak()

	p1 := person{
		first: "Dr.",
		last: "yes",
	}

	bar(p1)
	bar(sa1)
	bar(sa2)

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}