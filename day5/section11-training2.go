package main

import "fmt"

func inputNums() []int {
	var num int
	var nums []int

	for i := 0; i < 5; i++ {
		fmt.Scanln(&num)
		nums = append(nums, num)
	}

	return nums
}

func calExam(arr []int) (int, int, int) {
	var over []int
	var low []int
	var result int
	for _, value := range arr {
		if value >= 90 {
			over = append(over, value)
		}
		if value < 70 {
			low = append(low, value)
		}
		result += value
	}
	return result, len(over), len(low)
}

func printResult(sum int, over int, low int) {
	var result bool = true

	if sum < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	}
	if over < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	}
	if low > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}

	if result {
		fmt.Println("아이패드를 살 수 있습니다.")
	} else {
		fmt.Println("아이패드를 살 수 없습니다.")
	}
}

func main() {
	nums := inputNums()
	printResult(calExam(nums))
}
