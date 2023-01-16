package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the outer loop: %d\n", i)
		for j:= 0; j < 3; j++ {
			// fmt.Printf("the outer loop: %d\t the inner loop: %d\n", i, j)
			fmt.Printf("\t\tthe inner loop: %d\n", j)
		}
		
	}
}