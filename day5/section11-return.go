package main

import "fmt"

func add(num ...int) (int, int) {
	var result int
	var count int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
		count++
	}

	return result, count
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println(add(nums...))
}
