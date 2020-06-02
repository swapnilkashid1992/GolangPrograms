/*Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors. */

package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "swapnil",
		last:     "kashid",
		icecream: []string{"mango", "chocolet", "anjeer"},
	}
	p2 := person{
		first:    "aman",
		last:     "jagushte",
		icecream: []string{"wine", "vanilaa", "green apple"},
	}
	fmt.Println(p1, p2)
	for _, v := range p1.icecream {
		fmt.Println(v)
	}
	for _, v := range p2.icecream {
		fmt.Println(v)
	}
}
