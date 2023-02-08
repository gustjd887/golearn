package main

import "fmt"

func main() {
	x := factorial(4)
	fmt.Println(x)
	fmt.Println(loopFactorial(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	for i := n - 1; i > 0; i-- {
		n *= i
	}
	return n
}
