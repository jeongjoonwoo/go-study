package main

import "fmt"

var d int = 100

func exampleFunc1() {
	d++
	fmt.Println("d값", d)
}

func exampleFunc2() {
	var d int = 10
	d++
	fmt.Println("d값", d)
}

func main() {
	exampleFunc1()
	exampleFunc2()
	fmt.Println("main의 d는", d)
}
