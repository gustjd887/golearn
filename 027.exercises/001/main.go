package main

import (
	"fmt"

	"github.com/gustjd887/golearn/027.exercises/001/dog"
)

func main() {
	hy := 33
	dy := dog.Year(hy)
	fmt.Printf("human : %d, dog : %d\n", hy, dy)
}
