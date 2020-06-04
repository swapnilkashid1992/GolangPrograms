package main

import "fmt"

func main() {
	x := []int{10, 23, 45, 67, 8, 23, 44, 55, 68, 34, 56, 877}
	fmt.Println("After Add AT End")
	y := addAtEnd(&x, 99)
	fmt.Println(*y)
	fmt.Println("After Add AT First")
	y = addAtFirst(&x, 105)
	fmt.Println(*y)
	fmt.Println("After Delete From End")
	y = deleteFromEnd(&x)
	fmt.Println(*y)
	fmt.Println("After Delete from first")
	y = deleteFromFirst(&x)
	fmt.Println(*y)
}

func addAtEnd(x *[]int, num int) *[]int {
	y := *x
	y = append(y, num)
	*x = y
	return x
}
func addAtFirst(x *[]int, num int) *[]int {
	y := *x
	y = append([]int{num}, y...)
	*x = y
	return x
}

func deleteFromEnd(x *[]int) *[]int {
	y := *x
	if y == nil {
		fmt.Println("Queue is empty")
	}
	y = y[0 : len(y)-1]
	*x = y
	return x
}

func deleteFromFirst(x *[]int) *[]int {
	y := *x
	if y == nil {
		fmt.Println("Queue is empty")
	}
	y = y[1:]
	*x = y
	return x
}
