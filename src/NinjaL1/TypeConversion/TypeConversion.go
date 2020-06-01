package main

import (
	"fmt"
)

type newType int

var x newType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 45
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
