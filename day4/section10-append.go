package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 4}
	fmt.Println(a, "a의길이", len(a), "a의길이", cap(a))

	b := []int{5}
	fmt.Println(b, "b의길이", len(b), "b의길이", cap(b))

	c := append(a, 5)
	fmt.Println(c, "c의길이", len(c), "c의길이", cap(c))
}
