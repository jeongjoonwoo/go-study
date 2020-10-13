package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	}
	for _, num := range arr {
		fmt.Printf("값은 %d입니다.\n", num)
	}
	for index := range arr {
		fmt.Printf("arr[%d]입니다.\n", index)
	}
	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}

}
