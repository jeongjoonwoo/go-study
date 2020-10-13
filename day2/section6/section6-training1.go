package main

import "fmt"

func main() {
	var dan int
	fmt.Scan(&dan)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
	}
}
