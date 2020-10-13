package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64 // 둘레를 측정하는 메소드 추가
}

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

func (r Rect) perimeter() float64 { // 둘레를 측정하는 메소드 추가
	return 2 * (r.width + r.height)
}

func (c Circle) perimeter() float64 { // 둘레를 측정하는 메소드 추가
	return 2 * math.Pi * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

	printMeasure2(r1, c1, r2, c2)
}

func printMeasure2(m ...geometry) {
	for _, val := range m {
		fmt.Println("")
		fmt.Println(val.area())
		fmt.Println(val.perimeter())
	}
}
