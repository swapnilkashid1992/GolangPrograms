package main

import "fmt"

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	mslice := [][]string{x1, x2}
	fmt.Println(mslice)

	for i, v := range mslice {
		for j, v1 := range v {
			fmt.Printf("mslice[%d][%d]=%v\n", i, j, v1)
		}
	}
}
