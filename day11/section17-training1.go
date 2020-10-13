package main

import (
	"fmt"
	"sync"
)

func add(num1 int, num2 int, result *int, w *sync.WaitGroup) {
	defer w.Done()
	*result = num1 + num2
}

func main() {
	var num1, num2 int = 10, 5
	var result int
	var wait = new(sync.WaitGroup)

	wait.Add(1)

	go add(num1, num2, &result, wait)

	wait.Wait()

	fmt.Println(result)

}
