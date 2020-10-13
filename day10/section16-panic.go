package main

import "fmt"

func main() {
	var input string
	vals, err := fmt.Scanln(&input)
	fmt.Println("vals", vals)
	if err != nil {
		panic(err)
	}

	fmt.Println("에러가 nil일경우", input)
}
