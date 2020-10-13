# 고루틴(Goroutine)

## 비동기 프로세스의 기본

- Gorutin : 여러함수를 동시에 실행할 수 있는 논리적인 가상 스레드입니다.
  - 스레드 : 프로세스 안의 실행 흐름
  - 멀티스레드 : 여러가지 일을 동시에 처리가 가능하게 해주는 기술.
- 즉, Go에서 병렬 처리를 위해 사용되는것이 고루틴입니다.

```go
package main

import "fmt"

func testGo(){
    fmt.Println("Hello world!")
}

func main(){
    go testGo()

}
```

```

```

- `Hello world!` 가 `testGO`함수에서 실행이 되어야 하는데 실행하지 않고 종료됩니다.
- `testGo()`,`main()`이 동시에 실행되기 때문에, `testGo`실행되기 전에 `main()`함수가 종료됩니다.
- `testGo()`를 실행하기 위해서 `fmt.Scanln`을 이용하여 사용할 수 있습니다.
  - time package를 이용하여 잠시 main을 sleep시켜 출력할수도 있습니다.

```go
package main

import (
	"fmt"
	// "time"
)

func testGo() {
	fmt.Println("Hello world!")
}

func main() {
	go testGo()
    // time.Sleep(time.Second * 1)
    fmt.Scanln()
	fmt.Println("Seeya!")
}
```

```
//sleep을 하지않을경우
Hello world!
1
Seeya!
```

```
//sleep을 이용할경우
Hello world!
Seeya!
```

### 예제 - 난수생성

- 0<=n<30 순차적으로 실행하는 반복문으로 난수 생성해 함수 호출 대기시간을 설정.
- 비동기로 실행되어 숫자가 순차적으로 실행되지 않습니다.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
    r := rand.Intn(3) // 0부터 3까지 난수 생성
    fmt.Println(time.Duration(r))
	time.Sleep(time.Duration(r) * time.Second)
	// 난수를 Dration형으로 형변환 후 second로 계산
	fmt.Println(n)
}

func main() {
	for i := 0; i < 30; i++ {
		go hello(i)        // 고루틴 20개 생성 비동기 실행
	}

	fmt.Scanln()
}
```

```
0
16
20
7
5
10
29
24
13
28
15
18
6
12
11
1
26
4
27
21
9
2
17
19
14
3
22
8
23
25
end
```

## 고루틴의 활용

- 비동기로 실행되어 질때 main함수가 먼저 종료되는것을 막기위해 `fmt.Scanln()`을 이용했습니다.
- `sync` 패키지의 `WaitGroup`를 이용하여 고루틴이 끝날때 까지 대기하는 기능을 이용할 수 있습니다.

- `Add()` : 기다릴 고 루틴 수 설정
  - Done > Add : deadlock
  - Done < Add : Add의 수만큼 진행
- `Done()` : 고루틴이 실행된 함수 내에서 호출함으로써 함수 호출이 완료되었음을 알림
- `Wait()` : 고루틴이 모두 끝날때 까지 차단.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
    "sync"
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
```

- 앞의 난수생성하는 코드와 동일하지만. `fmt.Scanln`대신 `async`패키지를 이용했을 경우입니다.

1. `wait`변수에 new를 이용한 `WaitGroup`포인터 변수를 생성
   - call by referece를 위해 포인터를 이용해서 전달합니다.
2. `wait`에 30개의 고루틴 갯수 설정
3. `Done`를 이용하여 끝났음을 wait에 전달
4. `Wait()`는 async의 수만큼 진행되기 까지 대기

### 클로저에 고루틴

- `WaitGroup`는 클로저에 많이 이용됩니다.

```go
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
```

```
Hello
9
7
6
8
3
goorm!
2
4
0
1
5
```

- `str`변수를 선언하고 클로저에 직접 접근할 수 있습니다.

### 다중 CPU병렬처리

- 고루틴을 이용하면 실행되는 함수가 비동기 처리됩니다. 또한 많이 생성하여도 하나의 CPU에서 연산을 진행합니다.
- CPU의 코어가 보통 2개이상인 요즘에 CPU를 이용한 병렬처리를 지원합니다.
- 동시성 : 독립적으로 실행되는 기능
- 병렬처리 : 계산을 동시에 실행
- `runtime`패키지를 이용하면 간단하게 처리할 수 있습니다.
  1. `runtime.NumCPU()` : 현재 디바이스 CPU개수를 반환
  2. `runtime.GOMAXPROCS()` : 입력한 수만큼 CPU사용, 1미만일때 현재 설정값을 반환하고 설정값은 변하지 않습니다.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//디바이스의 총 CPU 개수를 반환하고 그 값을 CPU 사용 값으로 설정
	fmt.Println(runtime.GOMAXPROCS(0))
	// 현재 설정값 출력, 1미만이기 때문에 설정값 바꾸지 않음
	var wait sync.WaitGroup
	wait.Add(100)

	for i := 0; i<10; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
```

## 실습1 - 고루틴 실습

```go
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
```
