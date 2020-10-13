package main

import "fmt"

func main() {
	var str = "Hello Goorm!"
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(str, i)
		}

		done <- true //채널에 true를 송신함
	}()

	<-done //수신함으로써 대기를 끝냄
}
