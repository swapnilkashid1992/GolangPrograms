package main

import "fmt"

func main() {
	x := []int{10, 43, 23, 56, 43, 24, 56, 78}
	fmt.Println("Before")
	fmt.Println(x)
	y := pop(&x)
	fmt.Println("After POP")
	fmt.Println(*y)

	z := push(&x, 55)
	fmt.Println("After PUSH")
	fmt.Println(*z)

	fmt.Println("Pick : ", pick(x))

}

func pop(x *[]int) *[]int {
	if *x == nil {
		fmt.Println("Stack is empty")
	}
	y := *x
	y = y[:len(y)-1]
	*x = y
	return x
}

func push(x *[]int, y int) *[]int {
	z := *x
	*x = append(z, y)
	return x
}

func pick(x []int) int {
	return x[len(x)-1]
}
