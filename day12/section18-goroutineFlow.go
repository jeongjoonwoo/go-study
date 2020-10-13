package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int

	func() {
		result = a + b
	}()

	fmt.Printf("두 수의 합은 %d입니다.\n", result)
	result = 0
	go func() {
		result = a + b
	}()
	fmt.Printf("두 수의 합은 %d입니다.\n", result)
}
