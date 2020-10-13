package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("-로 숫자를 나누어 2개 입력하세요")
	fmt.Scanf("%d-%d", &x, &y)

	fmt.Println("첫숫자 ", x)
	fmt.Println("뒷숫자 ", y)
}
