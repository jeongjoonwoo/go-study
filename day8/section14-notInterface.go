package main

import (
	"fmt"
	"math" //Pi를 사용하기 위해 import함
)

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}

	fmt.Println(r1.area())
	fmt.Println(c1.area())
}
