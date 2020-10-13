package main

import "fmt"

func main() {
	var multiArray [2][3][4]int //3차원 배열 선언
	multiArray[1][1][2] = 10    // 인덱스를 이용한 값 초기화package main
	fmt.Println(multiArray)

	var a = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}, //3x3배열 초기화
	}

	fmt.Println(a[1][2]) //2행 3열의 값 출력
}
