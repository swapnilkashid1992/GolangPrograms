/*Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

package main

import "fmt"

func main(){
	a:= foo()
	b,c:= bar()
	fmt.Printf("Output of foo %v\n",a)
	fmt.Printf("Output of bar %v, %v",b,c)

}
func foo() int {
	x:=42
	fmt.Println("Inside Foo")
	return x 
}

func bar() (int,string) {
	x:=43
	y:="Swapnil"
	fmt.Println("Inside bar")
	return x,y
}
