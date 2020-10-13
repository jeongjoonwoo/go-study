package main

import "fmt"

func add(num1 int, num2 int, c chan int) {
	c <- num1 + num2
}

func main() {
	var num1, num2 int
	var c = make(chan int)

	fmt.Scanln(&num1, &num2)

	go add(num1, num2, c)

	result := <-c

	fmt.Println(result)
}
