package main

import "fmt"

func main() {
	fmt.Println("안녕하세요!")

	foo() //함수 호출
	fmt.Println("sequence flow") // sequence

	for i := 0; i < 100; i++ { // loop
		if i%2 == 0 { // conditional
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("나는 foo...")
}

func bar() {
	fmt.Println("and then we exited...")
}
