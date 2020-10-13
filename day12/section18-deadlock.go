package main

import "fmt"

func main() {
	c := make(chan int)

	func() {
		c <- 5
	}()

	fmt.Println(<-c)
}
