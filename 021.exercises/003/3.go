package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter = 0

	var wg sync.WaitGroup
	gs := 100
	wg.Add(gs)

	fmt.Println("start goroutine num", runtime.NumGoroutine())
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	fmt.Println("mid goroutine num", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end goroutine num", runtime.NumGoroutine())
	fmt.Println(counter)
}
