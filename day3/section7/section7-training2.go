package main

import "fmt"

func main() {
	var num1, num2, result int

	fmt.Scanf("%d %d", &num1, &num2)

	if num1 > num2 {
		result = num1 - num2
	} else {
		result = num2 - num1
	}

	fmt.Print(result)
}
