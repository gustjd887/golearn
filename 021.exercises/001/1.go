package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Start numCpu: ", runtime.NumCPU())
	fmt.Println("Start numGoroutine: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const wc = 2
	wg.Add(wc)
	fmt.Println("main go")

	go func() {
		fmt.Println("func1 go")
		runtime.Gosched()
		wg.Done()
	}()

	go func() {
		fmt.Println("func2 go")
		runtime.Gosched()
		wg.Done()
	}()
	fmt.Println("Mid numCpu: ", runtime.NumCPU())
	fmt.Println("Mid numGoroutine: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("End numCpu: ", runtime.NumCPU())
	fmt.Println("End numGoroutine: ", runtime.NumGoroutine())
}
