package main

import "fmt"

func main() {
	var fruit, color string

	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)

	switch fruit {
	case "apple":
		color = "red"
	case "banana":
		color = "yello"
	case "grape":
		color = "purple"
	default:
		fmt.Println("모르겠습니다")
		return
	}
	fmt.Println(color)
}
