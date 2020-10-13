package main

import "fmt"

func main() {
	var num int
	fmt.Print("정수 입력:")
	fmt.Scanln(&num)

	switch num {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	default:
		fmt.Println("모르겠어요.")
	}
}
