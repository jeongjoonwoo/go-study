package main

import "fmt"

func main() {
	var result int

	for num1 := 0; num1 < 10; num1++ {
		for num2 := 0; num2 < 10; num2++ {
			result = num1*10 + num2 + num2*10 + num1
			if result == 99 {
				fmt.Printf("%d%d + %d%d = %d\n", num1, num2, num2, num1, result)
			}
		}
	}
}
