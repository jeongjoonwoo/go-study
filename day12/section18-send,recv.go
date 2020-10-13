package main

import "fmt"

func main() {
	ch := sum(10, 5)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sum(num1, num2 int) <-chan int {
	result := make(chan int)

	go func() {
		result <- num1 + num2
		result <- 10
	}()

	return result
}
