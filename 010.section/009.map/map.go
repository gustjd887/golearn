package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) // 없으면 0 출력하게 됨.

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the if print", v)
	}

}
