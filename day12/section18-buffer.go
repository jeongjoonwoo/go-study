package main

import "fmt"

func main() {
	c := make(chan int, 1)

	func() {
		c <- 5
	}()

	fmt.Println(<-c)
}
