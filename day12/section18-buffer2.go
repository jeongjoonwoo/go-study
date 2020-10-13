package main

import (
	"fmt"
)

func main() {
	done := make(chan bool, 2)

	go func() {
		for i := 0; i < 6; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 6; i++ {
		<-done

		fmt.Println("메인 함수 : ", i)
	}
}
