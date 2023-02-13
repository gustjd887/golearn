package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	// for v := range c1 {
	// 	wg.Add(1)
	// 	go func(v int) {
	// 		c2 <- timeConsumingWork(v)
	// 		wg.Done()
	// 	}(v)
	// 	fmt.Println("NumGoroutime: ", runtime.NumGoroutine())
	// }

	const goroutines = 10
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() { // 10개의 goroutine으로 나눠서 처리하기 때문에 조금 천천히 처리되는건가...?
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
		fmt.Println("NumGoroutime: ", runtime.NumGoroutine())
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
