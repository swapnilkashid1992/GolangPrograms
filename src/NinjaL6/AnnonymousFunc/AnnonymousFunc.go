/*Build and use an anonymous func
 */
package main

import "fmt"

func main() {
	//func() {
	//	fmt.Println("Called Annonymous func")
	//}

	a := func(b int) {
		fmt.Println("called a ", b)
	}
	a(10)
}
