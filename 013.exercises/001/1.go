package main

import "fmt"

type person struct {
	first string
	last  string
	ice   []string
}

func main() {
	p1 := person{
		first: "hs",
		last:  "p",
		ice: []string{
			"mint",
			"banana",
		},
	}
	p2 := person{
		first: "ms",
		last:  "p",
		ice: []string{
			"choco",
			"apple",
		},
	}

	fmt.Printf("%v %v\n", p1.first, p1.last)
	for i, v := range p1.ice {
		fmt.Println(i+1, v)
	}

	fmt.Println("")

	fmt.Printf("%v %v\n", p2.first, p2.last)
	for i, v := range p2.ice {
		fmt.Println(i+1, v)
	}
}
