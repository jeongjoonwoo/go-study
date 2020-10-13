package main

import "fmt"

func main() {
	a, b := 10, 20
	str := "Hello goorm!"

	result := func() int { // 익명함수 변수에 초기화
		b := 40
		a += b
		return a // main 함수 변수 바로 접근
	}()

	func() {
		fmt.Println(str) // main 함수 변수 바로 접근
	}()

	fmt.Println(result)
	fmt.Println(a, b)
}
