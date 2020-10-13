package main

import (
	"fmt"
	"time"
)

func testGo() {
	fmt.Println("Hello world!")
}

func main() {
	go testGo()
	time.Sleep(time.Second * 1)
	// fmt.Scanln()
	fmt.Println("Seeya!")
}
