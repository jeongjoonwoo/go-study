package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4}
	sliceB := []int{5, 6, 7, 8}

	sliceA = append(sliceA, sliceB...)
	fmt.Println("sliceA", sliceA)
	fmt.Println("sliceA", sliceB)
}
