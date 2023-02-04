package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for i, v := range x {
		fmt.Println("index is", i)
		for j, v2 := range v {
			fmt.Printf("%v %v\t", j, v2)
		}
		fmt.Println("")
	}

	x["phs"] = []string{"1", "2", "3"}

	for i, v := range x {
		fmt.Println("index is", i)
		for j, v2 := range v {
			fmt.Printf("%v %v\t", j, v2)
		}
		fmt.Println("")
	}

	delete(x, "no_dr")

	for i, v := range x {
		fmt.Println("index is", i)
		for j, v2 := range v {
			fmt.Printf("%v %v\t", j, v2)
		}
		fmt.Println("")
	}
}
