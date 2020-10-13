package main

import "fmt"

func main() {
	numsch := num(10, 5)
	result := sum(numsch)
	//채널 result는 수신만 할 수 있음
	fmt.Println(<-result)
}

func num(num1, num2 int) <-chan int {
	numch := make(chan int)

	go func() {
		numch <- num1
		numch <- num2 //송신 후
		close(numch)
	}()

	return numch //수신 전용 채널로 반환
}

func sum(c <-chan int) <-chan int {
	//채널 c는 수신만 할 수 있음
	sumch := make(chan int)

	go func() {
		r := 0
		for i := range c { //채널 c로부터 수신
			r = r + i
		}
		sumch <- r // 송신 후
	}()

	return sumch //수신 전용 채널로 반환
}
