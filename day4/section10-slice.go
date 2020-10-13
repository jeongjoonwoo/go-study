package main

import "fmt"

func main() {
	var a []int

	fmt.Println("a의값", a, "a의길이", len(a), "cap값", cap(a))
	a = []int{1, 2, 3}
	fmt.Println("a{1,2,3}의값", a, "a의길이", len(a), "cap값", cap(a))
	a = append(a, 4)
	fmt.Println("a append(4)의값", a, "a의길이", len(a), "cap값", cap(a))

	var b []int
	if b == nil {
		fmt.Println("b is nil")
	}
}
