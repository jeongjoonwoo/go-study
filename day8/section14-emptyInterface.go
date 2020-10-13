package main

import "fmt"

func printVal(i interface{}) {
	fmt.Println(i)
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func main() {
	var x interface{} //빈 인터페이스 선언

	x = 1
	printVal(x)

	x = "test"
	printVal(x)

	r1 := Rect{10, 20}
	x = r1
	printVal(x)
}
