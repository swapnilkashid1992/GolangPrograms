package main

import "fmt"

type newType int

var x newType = 42

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 45
	fmt.Println(x)
}
