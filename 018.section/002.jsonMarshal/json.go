package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []Person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
