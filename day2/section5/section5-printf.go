package main

import "fmt"

func main() {
	name, age := "Kim", 24
	fmt.Printf("안녕하세요 반갑습니다.")
	fmt.Printf("안녕하세요 반갑습니다.\n")

	fmt.Printf("안녕하세요 저는 %s 입니다. 나이는 %d살 입니다\n", name, age)

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d 배열 입니다\n", arr)
}
