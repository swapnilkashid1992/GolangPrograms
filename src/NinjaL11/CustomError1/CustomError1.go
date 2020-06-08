package main

import "fmt"

type CustError struct {
	err string
}

func (e CustError) Error() string {
	return fmt.Sprintf("here is the error: %v", e.err)
}

func main() {
	c1 := CustError{
		err: "Error in FOO",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
}
