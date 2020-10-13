package main

import "fmt"

func main() {
	fmt.Printf("공백 네 칸: %4d\n", 10)
	fmt.Printf("4자리중 10을 제외한 앞 2자리는 0으로 채움: %04d\n", 10)
	fmt.Printf("총 4자리중 각각 왼쪽에 정렬 됩니다.: %-4d%-4d\n", 10, 15)
	fmt.Printf("12.3456를 소수점 둘째 자리까지만 표시하면 %.2f입니다.\n", 12.3456)
}
