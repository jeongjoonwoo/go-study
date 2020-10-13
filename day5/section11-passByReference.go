package main

import "fmt"

func printSqure(a *int) {
	*a *= *a

	fmt.Println(*a)
}
func main() {
	a := 4 //지역변수 선언

	printSqure(&a) //참조를 위한 a의 주솟값을 매개변수로 전달

	fmt.Println(a)
}
