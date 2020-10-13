package main

import "fmt"

func main() {
	var opt int
	var num1, num2, result float32

	fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:")
	fmt.Scan(&opt)
	if opt != 1 && opt != 2 && opt != 3 && opt != 4 {
		panic("1, 2, 3, 4중에 하나만 입력해야합니다!")
	}
	fmt.Print("두 개의 실수 입력:")
	fmt.Scan(&num1, &num2)

	if opt == 1 {
		result = num1 + num2
	} else if opt == 2 {
		result = num1 - num2
	} else if opt == 3 {
		result = num1 * num2
	} else if opt == 4 {
		result = num1 / num2
	}

	fmt.Printf("결과: %f\n", result)
}
