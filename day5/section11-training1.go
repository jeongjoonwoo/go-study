package main

import "fmt"

func bubbleSort(num ...int) {
	var temp int
	for i := 0; i < len(num)-1; i++ { //for문을 이용한 num[i] 순차 접근
		for j := i + 1; j < len(num); j++ {
			temp = num[i]
			if num[j] < temp {
				num[i] = num[j]
				num[j] = temp
			}
		}
	}
	return
}

func inputNums() []int {
	var num int
	var data int
	var nums []int
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		fmt.Scanln(&data)
		nums = append(nums, data)
	}
	return nums
}

func outputName(num ...int) {
	for _, v := range num {
		fmt.Printf("%d ", v)
	}
	return
}

func main() {
	nums := inputNums()
	bubbleSort(nums...)
	outputName(nums...)
}
