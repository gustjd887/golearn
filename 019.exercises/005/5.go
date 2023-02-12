package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (u ByAge) Len() int {
	return len(u)
}
func (u ByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}
func (u ByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type ByName []user

func (u ByName) Len() int {
	return len(u)
}
func (u ByName) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}
func (u ByName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	// your code goes here
	sort.Sort(ByAge(users))
	fmt.Println("[Sorted by Age]")
	for _, v := range users {
		fmt.Println(v.Age, v.Last, v.First)
		for _, v := range v.Sayings {
			fmt.Printf("\t%v\n", v)
		}
	}
	fmt.Println("[Sorted by Name]")
	sort.Sort(ByName(users))
	for _, v := range users {
		fmt.Println(v.Last, v.First, v.Age)
		for _, v := range v.Sayings {
			fmt.Printf("\t%v\n", v)
		}
	}
}
