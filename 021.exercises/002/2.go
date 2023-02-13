package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println(p.last, p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"hs", "p"}
	saySomething(&p1)
	// saySomething(p1) // 파라미터로 human을 받는 interface를 구성해 두었지만 호출 불가능.
	p1.speak() // 하지만 *person 포인터를 리시버로 받는 method는 호출 가능
}
