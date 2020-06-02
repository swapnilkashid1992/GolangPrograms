package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_jame":       []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(x)
	for k, v := range x {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Printf("\t\t%d\t%v\n", i, v1)
		}
	}
	x["swapnil"] = []string{"Movie", "Drama", "cooking"}
	for k, v := range x {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Printf("\t\t%d\t%v\n", i, v1)
		}
	}
	x = delete("swapnil")
	for k, v := range x {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Printf("\t\t%d\t%v\n", i, v1)
		}
	}
}
