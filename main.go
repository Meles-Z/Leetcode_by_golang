package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Reactangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (rect Reactangle) Area() float64 {
	return rect.height * rect.width
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func main() {
	var s Shape
	s = Reactangle{width: 3, height: 3}
	fmt.Println(s.(Reactangle).Area())
	s = Circle{radius: 3}
	fmt.Println(s.(Circle).Area())
}
