/*Create a func which returns a func
assign the returned func to a variable
call the returned func
*/
package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a())
}
func foo() func() int {
	return func() int {
		return 19
	}
}
