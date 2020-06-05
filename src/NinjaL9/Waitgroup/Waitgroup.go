package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go soo()
	wg.Wait()
	bar()

}

func foo() {
	fmt.Println("Inside Foo")
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}
	defer wg.Done()
}
func bar() {
	fmt.Println("Inside Bar")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
func soo() {
	fmt.Println("Inside soo")
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}
	defer wg.Done()
}
