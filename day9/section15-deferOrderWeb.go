package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
