package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int

	fmt.Print("")
	fmt.Scanln(&num1, &num2, &num3)

	var num4 = float32(num1)
	var num5 = uint(num2)
	var num6 = int64(num3)

	fmt.Printf("%T, %f\n%T, %d\n%T, %d\n", num4, num4, num5, num5, num6, num6)
}
