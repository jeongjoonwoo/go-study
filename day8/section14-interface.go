package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
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

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}
	fmt.Println(r1.area())
	fmt.Println(c1.area())
	printMeasure1(r1, c1, r2, c2)
}

func printMeasure1(m ...geometry) { //인터페이스를 가변 인자로 하는 함수
	for _, val := range m { //가변 인자 함수의 값은 슬라이스형
		fmt.Println(val.area()) //인터페이스의 메소드 호출
	}
}
