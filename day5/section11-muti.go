package main

import "fmt"

/*기능들만 모아놓은 함수들*/
func guide() { //매개변수 X 반환 값 X
	fmt.Println("두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.\n두 정수를 차례로 띄어 써주세요.")
	fmt.Print("두 정수를 입력해주세요 :")
}

func input() (int, int) { //매개변수 X 반환 값 O(두 개)
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func multi(a, b int) int { //매개변수 O, 반환 값 O
	return a * b
}

func printResult(num int) { //매개변수 O, 반환 값 X
	fmt.Printf("결과값은 %d입니다. 프로그램을 종료합니다.\n", num)
}

func main() { //간결해진 main문
	guide()
	num1, num2 := input()
	result := multi(num1, num2)
	printResult(result)
}
