package main

import "fmt"

func main() {
	//지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "충청남도"
	m["053"] = "대구광역시"

	fmt.Println(m)

	//동일한 key값으로 value값을 저장하면 갱신이 된다
	m["032"] = "인천"

	fmt.Println(m)

	//m에 있는 "031"key의 value와 함께 삭제
	delete(m, "031")

	fmt.Println(m)
}
