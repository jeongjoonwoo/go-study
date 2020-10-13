package main

import "fmt"

func main() {
	var sum = 0
	var i int

	i = 1

	for {
		sum += i
		if sum > 100 {
			break
		}
		i++
	}

	fmt.Println("1에서 ", i, " 까지 더하면 처음으로 100이 넘어요.")
	fmt.Println("합계:", sum)
}
