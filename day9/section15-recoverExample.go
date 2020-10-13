package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			main()
		}
	}()

	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	result := num1 / num2

	fmt.Println(result)
}
