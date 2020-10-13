package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string
	fmt.Print("이름,성별,학교를 입력하세요")
	fmt.Scan(&name, &gen, &school)

	fmt.Println("이름", name, "성별", gen, "학교", school)
}
