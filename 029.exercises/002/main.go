package main

import (
	"fmt"

	"github.com/gustjd887/golearn/029.exercises/002/quote"
	"github.com/gustjd887/golearn/029.exercises/002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
