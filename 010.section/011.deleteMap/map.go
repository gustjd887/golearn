package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) // 없으면 0 출력하게 됨.

	m["phs"] = 33

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the if print", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{2, 3, 4, 5, 6}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "phs")
	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["phs"]; ok {
		fmt.Println("value", v)
	}
}
