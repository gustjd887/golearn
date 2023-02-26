package main

import (
	"fmt"
)

func main() {

	var x = 5

	fmt.Printf("x 이진수: [%b]\n", x) // x 이진수: [101]

	fmt.Printf("3 이진수: [%3b]\n", 3) // 3 이진수: [ 11]

	x |= 3 // 111

	fmt.Printf("X |= 3 이진수: [%3b]\n", x) // X 이진수: [111]

}
