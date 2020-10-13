package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000000000"
	fmt.Println(p1)

	p2 := person{"nam", 31, "01022220000"} // 필드 이름을 생력할 시 순서대로 저장함
	fmt.Println(p2)

	p3 := person{contact: "01011110000", name: "park", age: 23} // 필드 이름을 명시할 시 순서와 상관 없이 저장할 수 있음
	fmt.Println(p3)

	p3.name = "ryu" //필드에 저장된 값을 수정할 수 있음
	fmt.Println(p3)

	fmt.Println(p3.contact) //필드 값의 개별 접근도 가능함
}
