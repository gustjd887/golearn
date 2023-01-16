package main

import "fmt"

func main () {
	s := "가"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // 문자열을 byte 리스트로 conversion
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println()
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}