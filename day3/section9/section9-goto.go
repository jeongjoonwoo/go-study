package main

import "fmt"

func main() {
	var num int

	fmt.Print("자연수를 입력하세요:")
	fmt.Scanln(&num)

	if num == 1 {
		goto ONE
	} else if num == 2 {
		goto TWO
	} else {
		goto OTHER
	}

ONE:
	fmt.Print("1을 입력했습니다.\n")
	goto END
TWO:
	fmt.Print("2를 입력했습니다.\n")
	goto END
OTHER:
	fmt.Print("1이나 2를 입력하지 않으셨습니다!\n")
END:
}
