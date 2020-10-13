package main

import "fmt"

func main() {
	var num1, num2 int = 3, 4

	var result float32 = float32(num1) / float32(num2)

	fmt.Printf("%f", result)
}
