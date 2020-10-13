package main

import "fmt"

func main() {

	for num1 := 2; num1 < 10; num1++ {
		if num1%2 == 0 {
			continue
		}
		for dan := 1; dan < 10; dan++ {
			if dan > num1 {
				continue
			}
			fmt.Printf("%d x %d = %d\n", num1, dan, num1*dan)
		}
		fmt.Printf("\n")
	}
}
