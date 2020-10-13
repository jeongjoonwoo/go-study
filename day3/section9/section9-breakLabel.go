package main

import "fmt"

func main() {
	var i int = 10

TEST1:
	for {
		fmt.Println(i)
		i -= 1
		if i == 0 {
			break TEST1
		}
	}

	fmt.Println("End")
}
