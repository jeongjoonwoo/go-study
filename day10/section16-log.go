package main

import (
	"fmt"
	"log"
)

func errorChek(n int) (string, error) {
	if n == 1 {
		s := "Goorm"
		return s, nil // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러 리턴
}

func main() {
	s, err := errorChek(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s) // Hello 1

	s, err = errorChek(2) // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Print(err)
	}
	fmt.Println(s)

	defer func() {
		s, err = errorChek(4) // 매개변수에 4를 넣었으므로 에러 발생
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	}()

	s, err = errorChek(3) // 매개변수에 3을 넣었으므로 에러 발생
	if err != nil {
		log.Panic(err) // defer 함수로 이동
	}
	fmt.Println(s)

	// 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
