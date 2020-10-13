package main

import "fmt"

func main() {
	var arr1 [5]int   //길이가 5인 int형 배열 arr1을 선언
	fmt.Println(arr1) //숫자를 선언하지 않고 출력해보기

	arr1 = [5]int{1, 2, 3, 4, 5}        //배열 초기화
	fmt.Println(arr1, arr1[0], arr1[4]) //배열 전체와 인덱스에 저장된 값들 출력해보기

	arr2 := [4]int{4, 5, 6, 7} //:= 를 이용해 선언
	arr2[0] = 32               //인덱스를 이용해 값을 초기화
	fmt.Println(arr2)          //arr2 전체 출력해보기

	var arr3 = [...]int{9, 8, 7, 6} //[...]을 이용한 배열 크기 자동 설정
	fmt.Println(arr3, len(arr3))    //arr3 전체와  배열 크기 출력해보기
}
