package main

import "fmt"

func main() {
	a := make([]int, 2, 3)

	for i := 0; i <= 10; i++ {
		a = append(a, i)
		fmt.Println(i, "값 추가시 길이 : ", len(a), "용량 크기 : ", cap(a))
	}
	fmt.Println(a)
}
