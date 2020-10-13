package main

import "fmt"

func printSqure(a int) {
	a *= a

	fmt.Println(a)
}
func main() {
	a := 4 //지역변수 선언

	printSqure(a)

	fmt.Println(a)
}
