package main

import "fmt"

func main() {
	c := make(chan string, 2) // 버퍼 2개 생성

	// 채널(버퍼)에 송신
	c <- "Hello"
	c <- "goorm"

	close(c) // 채널 닫음

	// 채널 수신
	val, open := <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open) // 무한 대기 상황 발생 x
	val, open = <-c
	fmt.Println(val, open)
}
