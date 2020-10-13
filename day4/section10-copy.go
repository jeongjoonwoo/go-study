package main

import "fmt"

func main() {
	sliceA := []int{0, 1, 2}
	sliceB := make([]int, 2, 4) //sliceA에 2배 용량인 슬라이스 선언
	sliceC := make([]int, 2, 4) //sliceA에 2배 용량인 슬라이스 선언

	copy(sliceB, sliceA) //A를 B에 붙여넣는다
	sliceB[1] = 10

	fmt.Println("sliceA : ", sliceA, "의 길이 :", len(sliceA), "용량 : ", cap(sliceA))
	fmt.Println("sliceB : ", sliceB, "의 길이 :", len(sliceB), "용량 : ", cap(sliceB))

	sliceC = sliceA[0:len(sliceA)]

	fmt.Println("sliceC : ", sliceC, "의 길이 :", len(sliceC), "용량 : ", cap(sliceC))

	fmt.Println("sliceA 1번 인덱스부터 출력:", sliceA[1:])
	fmt.Println("sliceA 1번 인덱스까지 출력:", sliceA[:2])

}
