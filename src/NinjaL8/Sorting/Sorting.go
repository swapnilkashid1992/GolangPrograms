/*sort int and string*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 4, 2, 5, 3, 6, 7, 6, 5, 8, 3, 9}
	b := []string{"Swapnil", "Aman", "Vinod", "Saatick", "Devendra"}
	sort.Ints(a)
	sort.Strings(b)
	fmt.Println(a)
	fmt.Println(b)

}
