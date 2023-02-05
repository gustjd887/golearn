package main

import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

func main() {
	p1 := struct { // anonymous struct, 하나의 작은 영역에서만 사용할 구조체가 필요할 경우 이렇게 사용
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
