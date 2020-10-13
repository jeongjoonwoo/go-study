package main

import "fmt"

func main() {
	var a, b int = 10, 0

	defer fmt.Println("Done")
	fmt.Println("Done1")
	result := a / b
	fmt.Println(result)
	fmt.Println("Done2")
}
