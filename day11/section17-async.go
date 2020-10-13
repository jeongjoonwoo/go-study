package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(n int, w *sync.WaitGroup) {
	defer w.Done() //끝났음을 전달

	r := rand.Intn(3)

	time.Sleep(time.Duration(r) * time.Second)

	fmt.Println(n)
}

func main() {
	wait := new(sync.WaitGroup) //waitgroup 생성

	wait.Add(30) // 100개의 고루틴을 기다림

	for i := 0; i < 30; i++ {
		go hello(i, wait) //wait을 매개변수로 전달
	}

	wait.Wait() // 고루틴이 모두 끝날때까지 대기
}
