/*create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

package main

import "fmt"

type SQUARE struct {
	length int
}
type CIRCLE struct {
	radius int
}

func (s SQUARE) AREA() float64 {
	return float64(s.length * s.length)
}
func (c CIRCLE) AREA() float64 {
	return 3.14 * float64(c.radius*c.radius)
}

type SHAPE interface {
	AREA() float64
}

func info(s SHAPE) {
	fmt.Printf("Area  is %v\n", s.AREA())
}

func main() {
	s1 := SQUARE{

		3,
	}
	c1 := CIRCLE{

		4,
	}
	info(s1)
	info(c1)

}
