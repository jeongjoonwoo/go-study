package main

import "fmt"

type triangle struct {
	width, height float32
}

func (s triangle) triAreaVal() float32 { //Value Receiver
	s.width += 10
	return s.width * s.height / 2
}

func (s *triangle) triAreaRef() float32 { //Pointer Reciver
	s.width += 10
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea_val := tri1.triAreaVal()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea_val)

	triarea_ref := tri1.triAreaRef()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea_ref)
}
