package main

import "fmt"

func main() {
	x := 20
	fmt.Printf("%d\t%b\t%x\n", x, x, x)
	x = x << 1
	fmt.Printf("%d\t%b\t%x", x, x, x)
}
