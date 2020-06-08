package main

import (
	"C:\\Users\\gs-1454\\Desktop\\GolagWorkspace\\GolangPrograms\\src\\Ninja12\\Test1\\dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
