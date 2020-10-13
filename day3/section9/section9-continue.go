package main

import "fmt"

func main() {
	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			fmt.Printf("%d", i)
		}

		fmt.Printf("%d ", i)
	}
	fmt.Print("\n")
	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
			continue //반복문 처음 부분으로 간다
		}

		fmt.Printf("%d ", i)
	}
}
