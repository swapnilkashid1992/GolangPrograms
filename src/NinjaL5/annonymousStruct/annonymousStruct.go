package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
	}{
		first: "Swapnil",
		last:  "kashid",
	}
	fmt.Println(p1)
}
