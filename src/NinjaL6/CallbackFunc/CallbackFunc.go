/*A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/

package main

import "fmt"

func main() {
	g := func(x []int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum is ", foo(g, x))
}
func foo(f func(x []int) int, y []int) int {
	return f(y)

}
