package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(12)

	str := "goorm!"

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func() {
		defer wait.Done()
		fmt.Println(str)
	}()

	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wait.Done()

			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
