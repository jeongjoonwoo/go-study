package main

import "fmt"

func main() {
	var a map[int]string

	if a == nil {
		fmt.Println("nil map")
	}

	var m = map[string]string{ //key:value, 형식으로 초기화한다
		"apple":  "red",
		"grape":  "purple",
		"banana": "yellow",
	}

	fmt.Println(m, "\nm의 길이는", len(m))
}
