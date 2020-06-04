package main

import (
	"fmt"
	"sort"
)

type User struct {
	Username1 string
	Password  string
}

type ByUsername1 []User

func (a ByUsername1) Len() int { return len(a) }
func (a ByUsername1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByUsername1) Less(i, j int) bool { return a[i].Username1 < a[j].Username1 }

func main() {
	u1 := User{
		Username1: "swapnil",
		Password:  "123456",
	}
	u2 := User{
		Username1: "Aman",
		Password:  "123456",
	}
	u3 := User{
		Username1: "Devendra",
		Password:  "123456",
	}
	u4 := User{
		Username1: "Madhura",
		Password:  "123456",
	}
	u5 := []User{u1, u2, u3, u4}
	sort.Sort(ByUsername1(u5))
	fmt.Println(u5)
}
