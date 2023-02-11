package main

import "fmt"

func main() {
	c := circle{5.1}
	info(c) // 주소를 전달해야 함

}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() { //reciver로 pointer를 받음
	fmt.Println(3.14 * c.radius * c.radius)
}

func (s square) area() {
	fmt.Println(s.length * s.length)
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}
