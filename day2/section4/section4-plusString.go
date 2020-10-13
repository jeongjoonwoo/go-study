package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	// Interpreted String Literal
	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}
