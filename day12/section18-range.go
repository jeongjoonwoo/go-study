package main

import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	for {
		if val, open := <-c; open { // 표현식; 조건식 형태
			// open이 true면 실행
			fmt.Println(val, open)
		} else {
			break
		}
	}
}
