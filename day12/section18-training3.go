package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			c <- true
		}
		fmt.Print("송신루틴완료")
	}()

	for i := 1; i <= 20; i++ {
		fmt.Println("수신한 데이터 :", i)
	}

}
