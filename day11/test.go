package main

import (
	"fmt"
	"math/rand"
)

func randomNumber(i int) int {
	r := rand.Intn(10) // 0부터 3까지 난수 생성
	return r
}

func main() {
	c := make([]int, 4)
	var b = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 4; i++ {
		c[i] = randomNumber(i, b)
	}

	fmt.Println(c)
}
