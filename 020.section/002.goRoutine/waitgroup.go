package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo() //cuncurrency 코드는 있지만, parallel 코드는 없기 때문에 순서대로 실행됨.
	bar()

	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
