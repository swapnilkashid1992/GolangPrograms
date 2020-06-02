/*Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person
then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.*/
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
	x := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range x {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, v1 := range v.icecream {
			fmt.Println(v1)
		}
		fmt.Println("--------------------------")
	}

}
