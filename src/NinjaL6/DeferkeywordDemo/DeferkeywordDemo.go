/*Use the “defer” keyword to show that a deferred func runs after the func containing it exits.*/
package main

import "fmt"

func main() {
	resourceOpen()
	defer resourceClosed()
	bar()
}
func resourceOpen() {
	fmt.Println("Resource open")
}
func bar() {
	fmt.Println("Inside bar")
}
func resourceClosed() {
	fmt.Println("Resource closed")
}
