package main

import "fmt"

func main() {
	jb := []string{"james", "Bond", "Chocolate", "martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
