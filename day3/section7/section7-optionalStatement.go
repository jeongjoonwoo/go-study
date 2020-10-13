package main

import "fmt"

func main() {

	num := 1
	val := 10
	if num == 1 { //괄호없이
		fmt.Println("num==1")
	}

	if val := num * 2; val == 2 { //괄호있이
		fmt.Println("val==2")
		fmt.Println("in if block val = ", val)
	}

	fmt.Println("out if block val = ", val)
}
