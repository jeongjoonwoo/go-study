package main

import "fmt"

func dessertList(fruit ...string) (count int, list []string) { //여기서 이미 선언된 것이다

	for i := 0; i < len(fruit); i++ {
		list = append(list, fruit[i])
		count++
	}

	return //생략하면 안 된다
}

func inputFruit() (list []string) { //Named return parameter는 값이 하나라도 괄호를 써야한다

	for {
		var fruit string
		fmt.Print("과일을 입력하세요:")
		fmt.Scanln(&fruit)

		if fruit != "1" {
			list = append(list, fruit)
		} else {
			fmt.Println("입력을 종료합니다.\n")
			break //반복문을 빠져나간다
		}
	}

	return
}

func main() {
	fmt.Println("디저트로 먹을 과일을 입력하고 출력합니다. \n1을 입력하면 입력을 멈춥니다.\n")
	count, list := dessertList(inputFruit()...) //함수를 변수처럼 사용할 수 있습니다
	fmt.Printf("%d개의 과일을 입력하셨고, 입력한 과일의 리스트는 %s입니다.\n", count, list)
}