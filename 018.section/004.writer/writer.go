package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Fprint(os.Stdout, "Hello")

}
