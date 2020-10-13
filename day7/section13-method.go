package main

import "fmt"

type triangle struct {
	width, height float32
}

func triArea(s *triangle) float32 { //'new'로 생성한 구조체 객체는 포인터값 반환
	return s.width * s.height / 2 //포인터 구조체는 자동 역참조 "*" 생략
}

func (s triangle) triArea() float32 { //value receiver
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea := triArea(tri1)
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea)

	triarea2 := tri1.triArea()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea2)
}
