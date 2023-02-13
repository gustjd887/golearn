package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	gs := 100
	wg.Add(gs)

	fmt.Println("start goroutine num", runtime.NumGoroutine())
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	fmt.Println("mid goroutine num", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end goroutine num", runtime.NumGoroutine())
	fmt.Println(counter)
}
