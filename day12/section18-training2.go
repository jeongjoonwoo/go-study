package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	timer := 11
	state := false
	go sendMessage(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		timer--
		select {
		case <-ch:
			fmt.Println("메세제지가 발송되었습니다.")
			state = true
		default:
			go func() {
				if timer == 0 {
					fmt.Println("메세지 발송에 실패했습니다.")
					state = true

				} else {
					fmt.Println(timer, "초 안에 메세지를 입력하시오.")
				}
			}()
		}
		if state {
			break
		}
	}
}

func sendMessage(ch chan string) {
	var message string
	fmt.Scanln(&message)
	ch <- message
}
