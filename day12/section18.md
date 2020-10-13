# 채널(Channel)

## 고루틴의 데이터 통로 : 채널

```go
//section18-goroutineFlow.go
package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int

	func() { //일반적인 익명함수
		result = a + b
	}()

    fmt.Printf("두 수의 합은 %d입니다.\n", result)

	result = 0
	go func() { //고루틴이 들어간 익명함수
		result = a + b
	}()
	fmt.Printf("두 수의 합은 %d입니다.\n", result)
}
```

```
두 수의 합은 15입니다.
두 수의 합은 0입니다.
```

- 익명함수 클로저에 두 수를 더한 값을 `result`에 넣는 프로그램입니다.
- 일반적인 익명함수
  - 일반적인 익명함수는 흐름대로 15의 값이 나오는것을 알 수 있습니다.
- Goroutine를 이용한 익명함수
  - go가 들어간 익명함수는 0의값이 나옵니다.
- 앞에서 배운 `fmt.Scan`,`sync.WaitGroup`는 모든 고루틴이 종료할때까지 기다리는 용법입니다.
  - 문제점 : 고루틴 사이의 흐름을 개별적으로 제어하지 않습니다.

#### 채널 : 프로그램의 흐름대로 진행될때, 채널은 생성되면 고루틴에 들어가서 흐름을 진행합니다.

- 채널은 고루틴 사이에서 값을 송/수신하는 통로역활입니다.
- **선언방식 : `make(chan 데이터타입)`**
- 송/수신 : `<-`
- 값 전달 : 채널`<-`데이터
- 값 호출 : `<-`채널
- 채널은 호출과 함께 함수가 종료되어 다른 변수에 대입을 하여 이용해야합니다.

```go
//section18-channelFlow.go
package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int

	c := make(chan int)

	go func() {
        c <- a + b
        c <- 1000
	}()

	result = <-c
	fmt.Printf("두 수의 합은 %d입니다.\n", result)
}
```

```
두 수의 합은 15입니다.
```

#### 일반적인 흐름 과 채널을 이용한 흐름 차이

##### 일반적인 흐름

1. a,b,result선언
2. 익명함수
   1. result = a+b 연산
3. result출력
4. 함수종료

##### 채널을 이용하지 않은 고루틴의 흐름

1. a,b,result선언
2. 고루틴 생성
   1. result = a+b 연산
   2. 함수종료
3. result출력
4. 함수종료

##### 채널을 이용한 고루틴의 흐름

1. a,b,result 선언
2. 채널 c 생성
3. 고루틴 생성
   1. 채널c에 a+b데이터 송신
   2. 채널c 수신완료까지 대기
4. c에 데이터가 수신하기 까지 대기
   1. 데이터 수신 완료와함께 고루틴 함수 종료
5. result변수에 c데이터 수신
6. result출력
7. 함수종료

- 채널은 변수에 넣는거 이외에도 `fmt.Print(<-c)`형식으로도 이용할 수 있습니다.

```go
//section18-channelEx.go
package main

import "fmt"

func main() {
	var str = "Hello Goorm!"
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(str, i)
		}

		done <- true //채널에 true를 송신함
	}()

	<- done //수신함으로써 대기를 끝냄
}
```

```
Hello Goorm! 0
Hello Goorm! 1
Hello Goorm! 2
Hello Goorm! 3
Hello Goorm! 4
Hello Goorm! 5
Hello Goorm! 6
Hello Goorm! 7
Hello Goorm! 8
Hello Goorm! 9
```

---

## 비동기 채널과 버퍼

### 데드락(deadlock)

- 채널에서 자주 발생하는 오류가 교착상태입니다.
- deadlock(교착상태) : 둘 이상의 프로세스가 서로 가진 한정된 자원을 요청하는 경우 발생합니다.
  - 프로세스가 진행하지 못하고 대기상태가 되는것을 의미합니다.

```go
//section18-deadlock.go
package main

import "fmt"

func main() {
c := make(chan int)

func() {
    c <- 5
}()

fmt.Println(<-c)
}
```

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main.func1(...)
        /Users/joonwoojeong/ojt/go-study/day12/aa.go:9
main.main()
        /Users/joonwoojeong/ojt/go-study/day12/aa.go:10 +0x5a
exit status 2
```

- 채널c는 데이터를 보내고 있는데 데이터를 받는 수신자가 없기 때문에 값을 수신할때까지 무한정 기다리는 데드락 현상이 발생합니다.

### 비동기 채널 버퍼

- 채널에서 송수신이 1:1로 대응하는 것을 볼 수 있습니다.
- 송신 루틴에서 수신루틴으로 바로 전달 => 특정 사이즈의 버퍼를 생성한 후 송수신은 버퍼만 합니다.
  - 버퍼 : A -> B로 데이터를 전송할때 임시로 데이터를 저장하는 공간
- **선언방식 : make(chan 데이터타입,버퍼개수)**
  - 버퍼갯수는 배열처럼 0부터 시작합니다.
- 송신 루틴은 수신자가 없어도 버퍼에 보내면 일을 끝내고, 수신 루틴은 값을 받으면 송신루틴이 끝나지 않아도 종료됩니다.

##### 채널을 이용한 고루틴의 흐름

1. a,b,result 선언
2. 채널 c 생성 => 버퍼 채널 생성
3. 고루틴 생성
   1. 채널c에 a+b데이터 송신 => 채널 버퍼에 데이터 송신
      - 버퍼가 가득 차면 대기
      - 버퍼에 빈공간이 생기면 프로세스 진행
   2. 채널c 수신완료까지 대기 => 버퍼가 가득 차면 대기
4. c에 데이터가 수신하기 까지 대기 => 버퍼에 값이 없으면 대기
   1. 데이터 수신 완료와함께 고루틴 함수 종료
5. result변수에 c데이터 수신
6. result출력
7. 함수종료

```go
//section18-buffer.go
package main

import "fmt"

func main() {
c := make(chan int,1)

func() {
    c <- 5
}()

fmt.Println(<-c)
}
```

```
5
```

- 앞에서 교착상태에 빠져 종료되었던 코드가 정상적으로 진행이 됩니다.

```go
//section18-buffer2.go
package main

import (
	"fmt"
)

func main() {
	done := make(chan bool, 2)

	go func() {
		for i := 0; i < 6; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 6; i++ {
		<-done

		fmt.Println("메인 함수 : ", i)
	}
}
```

```
// 결과값이 늘 동일하지 않음
고루틴 :  0
고루틴 :  1
고루틴 :  2
메인 함수 :  0
메인 함수 :  1
메인 함수 :  2
메인 함수 :  3 //수신이 완료되었기 때문에 기다리지 않고 진행
고루틴 :  3 //buffer가 비워져 값을 전송합니다.
고루틴 :  4
고루틴 :  5
메인 함수 :  4
메인 함수 :  5
```

- 버퍼사이즈 2의 채널 done이 생성한 후 6번 채널을 송/수신을 진행합니다.
- 버퍼에 값을 보내면 바로 다음 함수를 실행합니다.
- 버퍼가 가득차서 송신을 할 수 없을때는 보내지 않고 묶여버립니다.

---

## 동기채널

### 동기채널

- 비동기 채널은 송/수신 관계가 1:1을 맞추지 않으면 데드락에 걸립니다.
- 송/수신 채널이 여러개여서 송/수신루틴을 번갈아가며 실행합니다.

```go
//section18-syncChannel.go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for i := 0; i < 4; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 4; i++ {
		<-done

		fmt.Println("메인 함수 : ", i)

        // time.Sleep(time.Second)
	}
}
```

```
고루틴 :  0
메인 함수 :  0
고루틴 :  1
메인 함수 :  1
고루틴 :  2
메인 함수 :  2
고루틴 :  3
메인 함수 :  3
```

---

## 채널 닫기

### 채널 닫기(Close)

- 데드락이 걸릴 상황에서 close를 이용할 수 있습니다.
- 채널에 데이터를 송신한 후 채널이 닫혀도 수신은 할 수 있습니다.
- 선언방식 : **close(채널이름)**

```go
//section18-close.go
package main

import "fmt"

func main() {
	c := make(chan string, 2) // 버퍼 2개 생성

	// 채널(버퍼)에 송신
	c <- "Hello"
	c <- "goorm"

	// 채널 수신
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // 무한 대기 상황 발생 x
	fmt.Println(<-c)
}
```

```
Hello
goorm
deadlock ~~
```

- 송신은 2개인데 수신이 4개여서 데드락이 걸립니다.

```go
//section18-close.go
package main

import "fmt"

func main() {
	c := make(chan string, 2) // 버퍼 2개 생성

	// 채널(버퍼)에 송신
	c <- "Hello"
	c <- "goorm"

	close(c) // 채널 닫음

	// 채널 수신
	val, open := <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open) // 무한 대기 상황 발생 x
	val, open = <-c
	fmt.Println(val, open)
}
```

```
Hello true
goorm true
 false
 false
```

- 송신의 수 만큼 수신을 한 후 나머지 수신에는 nil값을 반환합니다.

1. 채널을 닫은 후 데이터를 채널에 송신하면 'send on closed channel'이라는 메세지와 함께 panic이 발생합니다.
2. 채널의 데이터를 모두 수신하면 수신 데이터가 nil값이 뜹니다.
3. `<-채널`은 `채널데이터`,`채널개폐여부` 총 2개의 값을 반환합니다.

### 채널 range문

- `for range`: 컬렉션의 갯수만큼 반복문을 실행합니다.
- 채널 `range`문 : 송신 데이터의 갯수만큼 반복문을 실행합니다.
- `close(채널)`를 이용하지 않으면 `range`를 계속 반복합니다.

```go
//section18-range.go
package main

import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i<10; i++ {
		c <- i
	}
	close(c)

	for {
		if val, open := <-c; open { // 표현식; 조건식 형태
			// open이 true면 실행
			fmt.Println(val, open)
		} else {
			break
		}
	}
}
```

```
0 true
1 true
2 true
3 true
4 true
5 true
6 true
7 true
8 true
9 true
```

```go
//section18-range2.go
package main

import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i<10; i++ {
		c <- i
	}
	close(c)

	for val := range c { // <- c를 사용하지 않음 open을 확인 x
		fmt.Println(val)
	}
}
```

```
0
1
2
3
4
5
6
7
8
9
```

- range1 : 채널의 계패여부에 따라서 열렸을때만 print문을 실행하고, 닫혀있으면 break로 for문을 종료합니다.
- range2 : 채널에 송신된 수 만큼 반복문 실행합니다

---

## 송신 전용,수신 전용 채널

### 송신/수신 전용채널

- 지금까지 함수는 송신은 송신만, 수신은 수신만 하였지만, 하나의 함수에서 송/수신을 모두 할 수 있습니다.
- 양방향으로 통신도 가능합니다.

```go
package main

import "fmt"

func main() {
	c := make(chan int)

	go channel1(c)
	go channel2(c)

	fmt.Scanln()
}

func channel1(ch chan int) {
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("done1")
}

func channel2(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- 3
	ch <- 4

	fmt.Println("done2")
}
```

```
1
2
3
4
done1
done2
```

- c는 버퍼가 없어서 동기채널입니다.
- `channel1`은 1,2를 넣고 `channel2`는 3,4를 넣습니다. 각각 출력하면 1,2,3,4가 출력됩니다.
- 채널을 함수의 매개변수로 전달하거나, 반환할때 채널로 송신만한지,수신만할지 선택할 수 있습니다.
- `chan <- 채널타입`,`<-chan 채널타입`

```go
//section18-channel.go
package main

import "fmt"

func main() {
	c := make(chan int)

	go sendChannle(c)
	go receiveChannel(c)

	fmt.Scanln()
}

func sendChannle(ch chan<- int) {
	ch <- 1
	ch <- 2
	fmt.Println("done1")
}

func receiveChannel(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("done2")

}
```

```
1
2
done2
done1

```

### 송/수신 채널의 활용

- 채널을 이용하기 위해서는 해당 루틴에 채널이 있어야 합니다.

```go
package main

import "fmt"

func main() {
	ch := sum(10, 5)

	fmt.Println(<-ch)
}

func sum(num1, num2 int) <-chan int {
	result := make(chan int)

	go func() {
		result <- num1 + num2
	}()

	return result
}
```

- `ch`변수는 sum에서 channel을 받습니다.

1. main함수 호출
   1. sum함수호출
   2. result채널 생성
   3. 고루틴 생성
   4. result채널 반환
2. `ch`내에 데이터 출력
   - 고루틴 종료

```go
package main

import "fmt"

func main() {
	numsch := num(10, 5)
	result := sum(numsch)
	//채널 result는 수신만 할 수 있음
	fmt.Println(<-result)
}

func num(num1, num2 int) <-chan int {
	numch := make(chan int)

	go func() {
		numch <- num1
		numch <- num2 //송신 후
		close(numch)
	}()

	return numch //수신 전용 채널로 반환
}

func sum(c <-chan int) <-chan int {
	//채널 c는 수신만 할 수 있음
	sumch := make(chan int)

	go func() {
		r := 0
		for i := range c { //채널 c로부터 수신
			r = r + i
		}
		sumch <- r // 송신 후
	}()

	return sumch //수신 전용 채널로 반환
}
```

---

## Select문

### 채널 select문

```go
//section18-select.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- true
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- true
		}
	}()

	go func() {
		for {
			<-ch1
			fmt.Println("ch1 수신")
			<-ch2
			fmt.Println("ch2 수신")
		}
	}()

	time.Sleep(5 * time.Second)
}
```

```
ch1 수신
ch2 수신
ch1 수신
ch2 수신
ch1 수신
ch2 수신
ch1 수신
ch2 수신
```

- 5초동안 sleep를 걸어놓은 상태에서 1,2를 출력하는 상태입니다.
- `ch1` : 0.5초마다 출력
- `ch2` : 1초마다 출력
- ch2는 ch1보다 먼저 출력이 되어도 ch1이 출력이 되어야 하기 때문에 기다리는 상태입니다.
- find.all을 쓴거처럼 가장 나중에 선택된 것이 출력이 됩니다.

```go
select {
case <- 채널1이름:
	//실행할 구문
case <- 채널2이름;
	//실행할 구문
	...
default:
	//모든 case의 채널에 데이터가 송신되지 않았을 때 실행
}
```

- `section18-select.go`의 방식대로 하면 비효율 적입니다.
  - ch2는 ch1보다 0.5초 빠르게하고 0.5초를 대기 하기 때문입니다.

```go
//section18-select2.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- true
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- true
		}
	}()

	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("ch1 수신")
			case <-ch2:
				fmt.Println("ch2 수신")
			}

		}
	}()

	time.Sleep(5 * time.Second)
}
```

```
ch2 수신
ch1 수신
ch2 수신
ch2 수신
ch1 수신
ch2 수신
ch2 수신
ch1 수신
ch2 수신
ch2 수신
ch1 수신
ch2 수신
ch2 수신
```

- `ch2`채널은 `ch1`과 달리 수신이 될때마다 출력이됩니다.

```go
//section18-select3.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20
		}
	}()

	go func() {
		for {
			select {
				case a := <-ch1:
					fmt.Printf("ch1 데이터 %d 수신\n", a)
				case b := <-ch2:
					fmt.Printf("ch2 데이터 %d 수신\n", b)
			}

		}
	}()

	time.Sleep(5 * time.Second)
}
```

```
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
```

- `switch~case`문을 이용하여 변수에 값을 초기화하여 사용할 수 있습니다.

1. `select`송신된 채널이 있을때 그 채널을 수신하는 기능
2. case를 이용해 채널에 데이터를 송신할 수 있습니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			c := <- ch3
			fmt.Printf("ch3 데이터 %s 수신\n", c)
		}
	}()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20
		}
	}()

	go func() {
		for {
			select {
				case a := <-ch1:
				fmt.Printf("ch1 데이터 %d 수신\n", a)
				case b := <-ch2:
				fmt.Printf("ch2 데이터 %d 수신\n", b)
				case ch3 <- "goorm":
				}
		}
	}()

	time.Sleep(5 * time.Second)
}
```

```
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
```

- `case ch3<-"goorm"` 부분을 보면 실행문이 없음에도 0.2초 단위로 출력이 되는것을 확인할 수 있습니다.

## 실습

### 실습1. 고루틴 실습2

```go
package main

import "fmt"

func add(num1 int, num2 int,c chan int) {
	c <- num1+num2
}

func main() {
	var num1, num2 int
	var c = make(chan int)

	fmt.Scanln(&num1, &num2)

	go add(num1,num2,c)

	result := <-c

	fmt.Println(result)
}
```

### 메세지 전송

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	timer := 11
	state := false
	go sendMessage(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		timer--
		select {
		case <-ch:
			fmt.Println("메세제지가 발송되었습니다.")
			state = true
		default:
			go func() {
				if timer == 0 {
					fmt.Println("메세지 발송에 실패했습니다.")
					state = true

				} else {
					fmt.Println(timer, "초 안에 메세지를 입력하시오.")
				}
			}()
		}
		if state {
			break
		}
	}
}

func sendMessage(ch chan string) {
	var message string
	fmt.Scanln(&message)
	ch <- message
}
```

### 실습3 동기채널 연습

```go
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)

	go func(){
		for i:=0 ; i<20;i++{
			c<-true
		}
		fmt.Print("송신루틴완료")
	}()


	for i:=1 ; i<=20;i++{
		fmt.Println("수신한 데이터 :",i)
	}

}
```

### 실습4 비동기 채널 실습

```go
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 50)

	go func() {
		for i := 0; i < 20; i++ {
			c <- true
		}
		fmt.Println("루틴 송신 완료")
	}()

	for i := 1; i <= 20; i++ {
		fmt.Println("수신한 데이터 :", i)
	}
}
```
