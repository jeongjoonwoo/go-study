package main

import "fmt"

func main() {

	for val := 1; val < 100; val++ {
		if val%7 == 0 || val%9 == 0 {
			fmt.Printf("%d ", val)
		}
	}
}
