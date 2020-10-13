package main

import (
	"fmt"
)

func inputSubNum() (int, error) {
	var num int
	fmt.Scanln(&num)
	if num > 0 {
		return num, nil
	}
	return 0, fmt.Errorf("잘못된 과목 수입니다.")

}

func average(num int) (float64, error) {
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 || score > 100 {
			return 0, fmt.Errorf("잘못된 점수입니다.")
		}
		total += score
	}

	avg := float64(total) / float64(num)
	return avg, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	subNum, subErr := inputSubNum()
	if subErr != nil {
		panic(subErr)
	}
	result, avgErr := average(subNum)
	if avgErr != nil {
		panic(avgErr)
	}

	fmt.Println(result)
}
