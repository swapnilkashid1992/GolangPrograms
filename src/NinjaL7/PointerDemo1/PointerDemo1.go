/*Create a value and assign it to a variable.
Print the address of that value.
*/

package main

import "fmt"

func main() {
	a := 43
	fmt.Println("value of a ", a)
	fmt.Println("Address of a ", &a)
	fmt.Println("Value of a ", *&a)
}
