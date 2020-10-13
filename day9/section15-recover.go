package main

import "fmt"

func panicTest() {

	var a = [4]int{1, 2, 3, 4}

	for i := 0; i < 10; i++ { //panic 발생
		fmt.Println(a[i])
	}
}

func main() {
	defer func() {
		r := recover() //복구 및 에러 메시지 초기화
		fmt.Println(r) //에러 메시지 출력
	}()
	panicTest()

	fmt.Println("Hello, world!") // panic이 발생했지만 계속 실행됨
}
