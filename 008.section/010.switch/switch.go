package main

import "fmt"

func main() {
	n := "bond"
	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case (2==4):
	// 	fmt.Println("this should not print2")
	switch n{
	case "money", "bond", "Do No":
		fmt.Println("money or bond or do no")
	case "what":
		fmt.Println("bond james")
	// // case (3==3):
	// // 	fmt.Println("print")
	// case (4==4):
	// 	fmt.Println("also true, does it print?")
	// 	fallthrough
	// case (4==5):
	// 	fmt.Println("not true")
	// 	fallthrough
	// case (3==3):
	// 	fmt.Println("print")
	// 	fallthrough
	default:
		fmt.Println("this is default")
	}
}