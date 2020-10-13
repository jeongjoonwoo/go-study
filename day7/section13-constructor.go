package main

import "fmt"

type mapStruct struct {
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}
	d.data = map[int]string{}
	return &d
}
func newStruct2() mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}
	d.data = map[int]string{}
	return d
}

func main() {
	s1 := newStruct() // 생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}
	s2.data = map[int]string{}
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)

	s3 := newStruct() // 생성자 호출
	s3.data[1] = "hello"
	s3.data[10] = "world"

	fmt.Println(s3.data[1])
}
