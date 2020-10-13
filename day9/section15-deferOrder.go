package main

import "fmt"

func plus(a int, b int) int {
	defer fmt.Println("hello")
	return a + b
}

func main() {
	//defer fmt.Println(a) 선언된지 않은 a라서 출력되지 않음
	a := 1
	b := 2
	a++
	{
		defer fmt.Println("defer", a)
		defer fmt.Println("defer2", plus(a, b))
	}
	fmt.Println(a)
	a++
	fmt.Println(a)
	defer fmt.Println("defer3", plus(a, b))
}
