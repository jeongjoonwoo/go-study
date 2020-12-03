package setting

import (
	"fmt"
)

//PrintMenu is show menu default Setting
func PrintMenu() {
	fmt.Println("1. 구매")
	fmt.Println("2. 잔여 수량 확인")
	fmt.Println("3. 잔여 마일리지 확인")
	fmt.Println("4. 배송 상태 확인")
	fmt.Println("5. 장바구니 확인")
	fmt.Println("6. 프로그램 종료")
	fmt.Println("------------------------")
}

//WaitChoice is done action befor what user to do
func WaitChoice() {
	fmt.Print("\n엔터를 입력하면 메뉴 화면으로 돌아갑니다.\n")
	fmt.Scanln()
}
