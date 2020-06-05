package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const ws = 100
	var counter int64
	wg.Add(ws)
	for i := 0; i < ws; i++ {

		go func() {
			atomic.AddInt64(&counter, 1)
			r := atomic.LoadInt64(&counter)
			fmt.Println(r)
			wg.Done()
		}()

		fmt.Println("CPUs:", runtime.NumCPU())
		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("count:", counter)
}
