# section15

## defer,panic,recover.

### defer

> defer: 함수 및 블럭이 종료되는 시점 이전에 실행

```java
try{
    //메모리 할당 및 구문실행
}catch{
    // 예외처리(논리적오류)
}finally{
    //마지막에 꼭 실행 및 할당 공간 반납
}
```

- java의 try~finallly같은 형식에서 finally같은 구문입니다.
- defer은 다른 언어와 다르게 블록,특정위치 의 조건이 필요하지 않습니다. 즉 변수처럼 함수앞에 사용하면 이용할 수 있습니다.
- 실행 분기에 따라 예외 처리가 많아질때 사용하기 좋습니다.
  - defer을 이용하면 흐름 중간에 에러가 발생하여도 마지막에 꼭 실행하고 프로그램을 종료하지 않기때문입니다.

```go
//section15-deferOrder.go
package main

import "fmt"

func plus(a int, b int) int {
	defer fmt.Println("hello")
	return a + b
}

func main() {
	//defer fmt.Println(a) 선언된지 않은 a라서 출력되지 않음
	a := 1
	b := 2

    a++ //2

    defer fmt.Println("defer", a) //2
    defer fmt.Println("defer2", plus(a, b)) //2+2

    fmt.Println(a) //2

	a++//3

    fmt.Println(a) //3

	defer fmt.Println("defer3", plus(a, b)) // 3+2
}
```

```
hello //defer2 plus defer
2 //a
3 //a
hello //defer3 plus defer
defer3 5 //defer3
defer2 4 //defer2
defer 2 //defer
```

1. 함수 블럭이 종료되기 직전에 defer은 바로 실행합니다.
2. defer의 실행위치에 따라서 변수의 값이 다릅니다.
3. defer은 스택처럼 함수가 종료가 되면 pop하듯이 진행됩니다.

### defer의 사용 예제

```go
package main

import	"fmt"

func main() {
	var a, b int = 10, 0
	defer fmt.Println("Done")
    fmt.Println("Done1")
	result := a / b
    fmt.Println(result)
    defer fmt.Println("Done2")
}
```

```
Done1
Done
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
        /Users/joonwoojeong/ojt/go-study/day9/section15-defer.go:9 +0x85
exit status 2
```

- Done가 실행되고 이후 오류가 출력이되어집니다.
- 파일을 열고 닫을때 이용됩니다.

### defer의 실행순서

```go
//section15-deferOrederWeb.go
package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()

	for i := 0; i <3; i++ {
		defer fmt.Println(i)
	}
}
```

```
Hello
2
1
0
world
```

main함수에서 world 실행 함수를 맨 처음 실행하였지만, 가장 마지막에 출력이 됩니다.
스택의 후입선출(LIFO)처럼 제일 나중에 선언한 defer이 먼저 출력됩니다.

### defer의 파일 예

- os패키지 : 파일 입/출력을 위한 패키지

```go
package main

import (
	"fmt"
	"os"
)

func Helloworld() {
	file, err := os.Open("test.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)

	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}

func main() {
	Helloworld()
	fmt.Println("Done")
}
```

```
안녕!
Done
```

1. `os.Open("test.txt")`를 이용하면 test.txt파일을 엽니다.
   - `file, err := os.Open("test.txt")` 파일,에러값으로 반환값이 두개입니다. `file`변수에는 파일,`err`변수에는 에러값을 초기화 합니다.
2. `if err != nil` 부분에서 만약 에러값을 넣어준다면, 해당 에러값이 출력이 최고 함수가 종료되고 아니면 출력이 되고 함수 종료
3. `buf` 변수에 byte 슬라이스를 생성 한 후 만약 에러가 있으면 에러를 출력하고 함수는 종료가 됩니다.
4. `buf`,`os.open`모두 통과하면, 비로소 text의값이 출력이 됩니다.

출력을 해주기 위해서는 파일이 변수에 담아놓아서 이용을 해야 하는 상황이고, 함수가 종료되면 close가 필요한 상황이지만,
defer을 이용하지 않으면 출력을 위하여 함수 마지막부분에 일반적으로 작성합니다.
만약, err가 발생하여서 에러가 검출되어 err값이 출력이 되면, 하단의 close함수를 실행하지 않고 종료가 되었을겁니다.

---

## 종료하는 Panic(),복구하는 recover()

### Panic

> panic: 문제가 없어 보이는 것을 에러를 발생시켜 프로그램을 종료하는 기능

- 오류 : 프로그램상 허용하지 않은 문법,비정상적인 상황
  - `var a int = 30.5`
- 예외 : 프로그램상 실행에는 문제가 없으나 논리적으로 문제가 되는 상황.
  - `fmt.Println(0/10)`

```go
//section15-exception.go
package main

import "fmt"

func main() {
	var a, b int = 10, 0

	result := a / b
	fmt.Println(result)
}
```

```
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
        /Users/joonwoojeong/ojt/go-study/day9/section15-panic.go:8 +0x11
exit status 2
```

```go
//section15-error.go
package main

import "fmt"

func main() {
	var num int = 30.5 //문법적인 오류
	fmt.Println(num)
}
```

```
# command-line-arguments
./section15-error.go:6:6: constant 30.5 truncated to integer
```

index 이상의 값을 입력하여 ,panic이 발생하지만 defer문의 panic이전에 선언되었기 때문에 정상적으로 출력이 되고 종료가 됩니다.

```go
package main

import "fmt"

func panicTest() {
	var a = [4]int{1,2,3,4}

	defer fmt.Println("Panic done")

	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	panicTest()

	fmt.Println("Hello, world!")
}
```

```
1
2
3
4
Panic done
panic: runtime error: index out of range [4] with length 4

goroutine 1 [running]:
main.panicTest()
        /Users/joonwoojeong/ojt/go-study/day9/section15-panic.go:11 +0x1ac
main.main()
        /Users/joonwoojeong/ojt/go-study/day9/section15-panic.go:16 +0x25
exit status 2
```

### 강제로 panic생성

```go
//section15-forcePanic.go
package main

import "fmt"

func main() {
    var opt int
    var num1, num2, result float32

    fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:")
    fmt.Scan(&opt)
	if opt != 1 && opt != 2 && opt != 3 && opt != 4 {
		panic("1, 2, 3, 4중에 하나만 입력해야합니다!")
	}
    fmt.Print("두 개의 실수 입력:")
    fmt.Scan(&num1, &num2)

    if opt == 1 {
        result = num1 + num2
    } else if opt == 2 {
        result = num1 - num2
    } else if opt == 3 {
        result = num1 * num2
    } else if opt == 4 {
        result = num1 / num2
    }

    fmt.Printf("결과: %f\n", result)
}
```

- panic이 존재할경우

```
1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:7
panic: 1, 2, 3, 4중에 하나만 입력해야합니다!

goroutine 1 [running]:
main.main()
        /Users/joonwoojeong/ojt/go-study/day9/section15-forcePanic.go:12 +0x370
exit status 2
```

- panic이 존재하지 않을경우

```
1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:7
두 개의 실수 입력:15 2
결과: 0.000000
```

개발자가 원하는 상황에 panic을 이용하여 프로그램을 종료 할 수 있습니다.

---

### Recover

- panic: 프로그램을 강제로 종료
- recover: panic상황에서 프로그램을 종료하지 않고 예외처리

```go
//section15-recover.go
package main

import "fmt"

func panicTest() {
	defer func() {
		r := recover() //복구 및 에러 메시지 초기화
		fmt.Println(r) //에러 메시지 출력
	}()

    var a = [4]int{1,2,3,4}

    for i := 0; i < 10; i++ { //panic 발생
        fmt.Println(a[i])
    }
}

func main() {
    panicTest()

    fmt.Println("Hello, world!") // panic이 발생했지만 계속 실행됨
}
```

```
1
2
3
4
runtime error: index out of range [4] with length 4
Hello, world!
```

- recover는 프로그램이 종료되지 않고, 복구하는 기능을합니다.
- 프로그램일 종료되기 직전에 실행해야 함으로 defer가 선언된 함수 안에서 쓰입니다.
- 에러 메세지를 반환합니다.
  - 변수에 초기화 하여서 에러 메레시를 출력할 수 있습니다.
  - 초기화 하지 않으면 에러 메세지를 출력하지 않습니다.

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println(r)

			main()
		}
	}()

	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	result := num1 / num2

	fmt.Println(result)
}
```

- recover 작성시

```
runtime error: integer divide by zero
3
runtime error: integer divide by zero
4
runtime error: integer divide by zero

runtime error: integer divide by zero
1
runtime error: integer divide by zero
3
runtime error: integer divide by zero
-20 103
0
```

- recover 없을시

```
0
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
        /Users/joonwoojeong/ojt/go-study/day9/section15-recoverExample.go:17 +0x17a
exit status 2
```

---

## 실습

### 실습 1. 입력

```go
package main

import "fmt"

func main() {
	var names[] string
	var name string

	for {
		fmt.Scanln(&name)
		if name == "0" {
			break
		} else {
			names= append(names,name)
		}
	}

	for _,val := range names {
		defer fmt.Println(val)
	}
}
```

### 실습 2. 중간고사 평균 점수 2

```go
//section15-training2.go
package main

import "fmt"

func average() float64 {
	var num int
	fmt.Scanln(&num)

	if num <= 0 {
		panic("잘못된 과목 수입니다.")
	}

	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)

		if score < 0 {
			panic("잘못된 점수입니다.")
		}
		total += score
	}

	avg := float64(total) / float64(num)

	return avg
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			main()
		}
	}()

	result := average()
	fmt.Println(result)
}

```
