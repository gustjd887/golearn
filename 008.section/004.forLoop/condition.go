package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 9 { // for condition 대신에 이런식으로 사용할 수도 있음
			break
		}
		fmt.Println(x)
		x++ // 증가를 주석처리하면 계속돌아감.
	}
	fmt.Printf("Done.")
}