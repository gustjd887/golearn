package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	// for i, v := range x {
	// 	fmt.Println("index is", i)
	// 	for j, v2 := range v {
	// 		fmt.Printf("%v %v\t", j, v2)
	// 	}
	// 	fmt.Println("")
	// } // 여기서 값만 출력하려면... 인덱스로 사용한 i나 j를 사용하지 않으려면 어떻게...?
	// range 대신 len을 사용하면 되나???
	// 한번 더 해보기

	for i := range x {
		fmt.Println(i)
		for j := range i {
			fmt.Printf("%v\t", j)
		}
		fmt.Println("")
	}
}
