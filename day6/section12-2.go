package main

import "fmt"

func next(count int) func(money int) int {
	i := count
	return func(money int) int {
		i += money
		// fmt.Println(&i)
		return i
	}
}

func main() {
	nextInt := next(10)
	fmt.Println(nextInt(10))
	fmt.Println(nextInt(10))
	fmt.Println(nextInt(10))

	// newInt := next(10)
	// fmt.Println(newInt(10))
	fmt.Println(nextInt(10))
}
