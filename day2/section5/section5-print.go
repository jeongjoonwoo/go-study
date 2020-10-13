package main

import "fmt"

func main() {
	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(n) //개행을 위해 \n를 입력한다.
	fmt.Println(arr)
	fmt.Println("여기까지 Println입니다.")

	fmt.Print(n) //개행을 위해 \n를 입력한다.
	fmt.Print(arr)
	fmt.Print("여기까지 Print입니다.")
}
