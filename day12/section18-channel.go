package main

import "fmt"

func main() {
	c := make(chan int)

	go sendChannle(c)
	go receiveChannel(c)

	fmt.Scanln()
}

func sendChannle(ch chan<- int) {
	ch <- 1
	ch <- 2
	fmt.Println("done1")
}

func receiveChannel(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("done2")

}
