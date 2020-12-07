package main

import (
	"fmt"
	pro "project/package/setting"
	"time"
)

type productInfo struct {
	price int
	count int
}

type productForm map[string]*productInfo

type buyer struct {
	mileage      int
	shoppingList map[string]int
}

type delivery struct {
	status   string
	stations map[string]int
}

func innerBuyer() *buyer {
	inner := buyer{}
	inner.mileage = 1000000
	inner.shoppingList = map[string]int{}
	return &inner
}

func newDelivery() delivery {
	dele := delivery{}
	dele.stations = map[string]int{}
	return dele
}

func main() {
	var selectedItem = 0
	var actionsNum = 0
	var numbuy = 0
	var deliveryList = make([]delivery, 5)
	var deliveryState = make(chan bool)
	var customer = *innerBuyer()
	var productList = []string{"tumbler", "longPadding", "backpack", "shoese", "snack"}
	var stock = productForm{
		"tumbler": {
			price: 10000,
			count: 30,
		},
		"longPadding": {
			price: 500000,
			count: 20,
		},
		"backpack": {
			price: 400000,
			count: 20,
		},
		"shoese": {
			price: 150000,
			count: 50,
		},
		"snack": {
			price: 1200,
			count: 500,
		},
	}
	for i := 0; i < 5; i++ {
		deliveryList[i] = newDelivery()
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		go deliveryStatus(deliveryState, i, deliveryList, &numbuy)
	}

	for {
		pro.PrintMenu()
		fmt.Print("실행할 기능을 입력하시오 : ")
		fmt.Scan(&actionsNum)
		fmt.Println("")

		if actionsNum == 1 {
			var product string
			for index, productName := range productList {
				producntInfo := stock[productName]
				fmt.Printf("%d.%s의 가격: %d, 잔여수량: %d\n", index+1, productName, producntInfo.price, producntInfo.count)
			}
			fmt.Println()
			fmt.Print("구매할 물품의 숫자를 선택하세요 : ")
			fmt.Scanln(&selectedItem)
			fmt.Println()

			switch selectedItem {
			case 1:
				product = "tumbler"
			case 2:
				product = "longPadding"
			case 3:
				product = "backpack"
			case 4:
				product = "shoese"
			case 5:
				product = "snack"
			default:
				break
			}
			productInfo := stock[product]
			buying(productInfo, &customer, product, &numbuy, deliveryState)
			pro.WaitChoice()
		} else if actionsNum == 2 {
			for index, productName := range productList {
				producntInfo := stock[productName]
				fmt.Printf("%d.%s 잔여수량: %d\n", index+1, productName, producntInfo.count)
			}
			pro.WaitChoice()
		} else if actionsNum == 3 {
			fmt.Println("현재 잔여 마일리지는 ", customer.mileage, "점 입니다")
			pro.WaitChoice()
		} else if actionsNum == 4 {
			for i := 0; i < len(deliveryList); i++ {
				fmt.Println(i+1, "번 배송상황:", deliveryList[i].status)
			}
			pro.WaitChoice()
		} else if actionsNum == 5 {
			totalPrice := 0
			fmt.Println("장바구니 목록")
			fmt.Println("------------------------")
			for productName, productCount := range customer.shoppingList {
				fmt.Printf("%s: %d\n", productName, productCount)
				totalPrice += productCount * stock[productName].price
			}
			fmt.Println("\n합계: ", totalPrice)
			cartList(totalPrice, &customer, &numbuy, deliveryState)
			fmt.Println("")
		} else if actionsNum == 6 {
			break
		} else {
			fmt.Println("잘못된 입력입니다.")
		}
	}
}

func buying(stock *productInfo, customer *buyer, product string, numbuy *int, deliveryState chan bool) {
	var productAmount = 0
	var action = 0
	fmt.Printf("선택한 제품 : %s\n구매하실 갯수를 선택해주세요 : ", product)
	fmt.Scanln(&productAmount)
	fmt.Println()
	if productAmount <= 0 {
		panic("올바른 수량을 입력하세요.")
	}

	if stock.count-productAmount < 0 {
		fmt.Println("구매하실 수 있는 갯수를 초과하였습니다.")
	} else {
		fmt.Printf("1. 바로주문 \n2. 장바구니 담기\n")
		fmt.Print("실행할 기능을 입력하시오 : ")
		fmt.Scanln(&action)
		if action == 1 && customer.mileage > productAmount*stock.price {
			if *numbuy < 5 {
				customer.mileage -= productAmount * stock.price
				stock.count -= productAmount
				deliveryState <- true
				fmt.Println("주문이 완료되었습니다.")
				fmt.Println("남은 마일리지: ", customer.mileage)
				*numbuy++
			} else {
				fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
			}
		} else if action == 2 {
			customer.shoppingList[product] += productAmount
			stock.count -= productAmount
		} else {
			fmt.Println("마일리지가 부족합니다.")
		}
	}
}

func cartList(totalPrice int, customer *buyer, numbuy *int, deliveryState chan bool) {
	var cartAction int
	if totalPrice != 0 {
		fmt.Println("1. 장바구니 상품 주문")
		fmt.Println("2. 메뉴로 돌아가기")
		fmt.Print("실행할 기능을 입력하시오 :")
		fmt.Scanln(&cartAction)
		fmt.Println()

		if cartAction == 1 {
			fmt.Println("보유 마일리지 : ", customer.mileage)
			fmt.Println("필요 마일리지 : ", totalPrice)
			fmt.Println()
			if totalPrice > customer.mileage {
				fmt.Println("마일리지가 부족합니다.")
			} else {
				if *numbuy < 5 {
					customer.mileage -= totalPrice
					deliveryState <- true
					fmt.Println("주문이 완료되었습니다.")
					fmt.Println("남은 마일리지: ", customer.mileage)
					*numbuy++
				} else {
					fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
				}
			}
		} else if cartAction == 2 {
			return
		} else {
			fmt.Println("잘못된 입력입니다.")
		}
	}
	fmt.Println("장바구니가 비었습니다")
}

func deliveryStatus(deliveryState chan bool, i int, deliveryList []delivery, numbuy *int) {
	for {
		if <-deliveryState {
			deliveryList[i].status = "주문접수"
			time.Sleep(time.Second * 10)

			deliveryList[i].status = "배송중"
			time.Sleep(time.Second * 30)

			deliveryList[i].status = "배송완료"
			time.Sleep(time.Second * 10)

			deliveryList[i].status = ""
			*numbuy--
		}
	}
}
