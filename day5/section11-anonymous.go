package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello")
	}()

	func(a int, b int) {
		result := a + b
		fmt.Println(result)
	}(1, 3)

	result := func(a string, b string) string {
		return a + b
	}("hello", " world!")
	fmt.Println(result)

	i, j := 10.2, 20.4
	divide := func(a float64, b float64) float64 {
		return i / j
	}(i, j)
	fmt.Println(divide)
}
