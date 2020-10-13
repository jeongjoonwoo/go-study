package main

import "fmt"

func main() {
	var fruit string

	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)

	if (fruit != "apple") && (fruit != "banana") && (fruit != "grape") {
		fmt.Println("잘못 입력했습니다.")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("RED")
	case "banana":
		fmt.Println("YELLOW")
	case "grape":
		fmt.Println("PURPLE")
	}
}
