package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const ws = 100
	counter := 0
	wg.Add(ws)
	for i := 0; i < ws; i++ {

		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
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
