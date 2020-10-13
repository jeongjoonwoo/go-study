package main

import "fmt"

func main() {
	//지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "충청남도"
	m["053"] = "대구광역시"

	val, exist := m["055"]
	fmt.Println("055 : ", val, exist)
	val, exist = m["02"]
	fmt.Println("02 : ", val, exist)
}
