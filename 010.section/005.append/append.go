package main

import "fmt"

func main() {
	x := []int{4,5,6,7,1}
	fmt.Println(len(x))
	x = append(x,1,2,3)
	fmt.Println(x)

	y := []int{2,4,1,2}
	x = append(x,y...) // y에 있는 모든 값들을 풀어쓸 때 이렇게 사용
	fmt.Println(x)
}