package main

import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	for val := range c { // <- c를 사용하지 않음
		fmt.Println(val)
	}
}
