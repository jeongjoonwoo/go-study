package main

import "fmt"

const ( 
	c1 = 10 //첫 번째 값은 무조건 초기화해야 함
	c2
	c3
	c4 = "goorm" //다른 형 선언 가능
	c5
	c6 = iota //c8까지 index값 저장
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}