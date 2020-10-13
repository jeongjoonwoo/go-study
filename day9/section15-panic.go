package main

import "fmt"

func panicTest() {
	var a = [4]int{1, 2, 3, 4}

	defer fmt.Println("Panic done")

	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	panicTest()

	fmt.Println("Hello, world!")
}
