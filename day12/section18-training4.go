package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
			fmt.Println("goro", i)
		}
		fmt.Println("루틴 송신 완료")
	}()

	for i := 1; i <= 20; i++ {

		fmt.Println("수신한 데이터 :", <-c)
	}
}
