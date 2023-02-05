package main

import "fmt"

type person struct {
	first string
	last  string
	ice   []string
}

func main() {
	// p1 := person{
	// 	first: "hs",
	// 	last:  "p",
	// 	ice: []string{
	// 		"mint",
	// 		"banana",
	// 	},
	// }

	m := map[string]person{
		"phs": {
			first: "hs",
			last:  "p",
			ice: []string{
				"mint",
				"banana",
			},
		},
		"pms": {
			first: "ms",
			last:  "p",
			ice: []string{
				"choco",
				"apple",
			},
		},
	}

	// p2 := person{
	// 	first: "ms",
	// 	last:  "p",
	// 	ice: []string{
	// 		"choco",
	// 		"apple",
	// 	},
	// }

	// fmt.Printf("%v %v\n", p1.first, p1.last)
	// for i, v := range p1.ice {
	// 	fmt.Println(i+1, v)
	// }

	// fmt.Println("")

	// fmt.Printf("%v %v\n", p2.first, p2.last)
	// for i, v := range p2.ice {
	// 	fmt.Println(i+1, v)
	// }
	fmt.Printf("%v %v\n", m["phs"].first, m["phs"].last)
	for i, v := range m["phs"].ice {
		fmt.Println(i+1, v)
	}
	fmt.Printf("%v %v\n", m["pms"].first, m["pms"].last)
	for i, v := range m["pms"].ice {
		fmt.Println(i+1, v)
	}

}
