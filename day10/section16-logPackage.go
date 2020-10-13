package main

import (
	"fmt"
	"log"
)

func divide(a float32, b float32) (result float32, err error) {
	if b == 0 {
		return 0, fmt.Errorf("%.2f으로 나누지마", b)
	}
	result = a / b
	return
}

func main() {
	var num1, num2 float32
	fmt.Scanln(&num1, &num2)

	result, err := divide(num1, num2)

	if err != nil {
		log.Print(err)
	}

	fmt.Println(result)
}
