package main

import "fmt"

const (
	x = 2020 + iota
	y = 2020 + iota
	z = 2020 + iota
	a = 2020 + iota
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
}
