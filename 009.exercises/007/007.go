package main

import "fmt"

func main() {
	i := "phs"

	if i == "phs" {
		fmt.Println("i is phs")
	} else if i == "hsp" {
		fmt.Println("i is hsp")
	} else {
		fmt.Printf("i is not phs or hsp")
	}
}