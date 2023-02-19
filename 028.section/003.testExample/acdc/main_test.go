package acdc

import (
	"fmt"
)

func ExampleSum() {
	fmt.Print(Sum(2, 3))
	// Output:
	// 5
}

func ExampleSum_second() {
	fmt.Print(Sum(2, 4))
	// Output:
	// 6
}
