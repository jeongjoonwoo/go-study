# 에러처리

## 에러처리의 기본

- 예외 처리 부분에서 문제가 발생 했을때, 앞에서는 프로그램이 종료가 되거나 정상적인 입력이 안되어서 문제가 생겼을때 별도의 처리를 해주지 않았습니다.
- 에러를 처리하는 별도의 부분이 존재하지 않기때문에, 컴파일러가 알아차리지 못하는 프로그램상의 오류를 예방하기 위해서 사용합니다.
- 반환값이 있는 경우 에러처러리를 결과값과 에러값을 함께 반환해야합니다.

```go
//section16-panic.go
package main

import "fmt"

func main() {
	var input int
	vals, err := fmt.Scanln(&input)
	fmt.Println("vals", vals)
	if err != nil {
		panic(err)
	}

	fmt.Println("에러가 nil일경우", input)
}
```

- Scanln도 반환값이 존재합니다.

  - vals : 타입에 맞는값을 입력하면, 1 아닐경우 0이 입력됩니다.
  - err : 타입에 맞는값을 입력하면 정상적으로 입력, 아닐경우 err가 출력됩니다.

- 값이 타입에 맞게 정상적으로 입력이 되면, 해당 값들이 정상적으로 입력이 `err`는 `nil`이 됩니다.

에러값 이용은 앞의 예외처리에서 정수에 0을 나누었을때, 연산이 되지 않아 `panic`상태가 되어 프로그램이 종료되기 때문에, 에러가 발생했을때, 정상적으로 에러처리를 찾고, 출력할 것인지 생각해주어야 합니다.

### 에러 값 설정

```go
type error interface{
    Error() string
}
```

- error
  - 타입 : 인터페이스
  - Error()이라는 string형을 반환하는 메소드 하나만 가지고 있습니다.

```go
func (e *errorString) Error() string{
    return e.s
}
```

- `Error`메소드의 원형에서 보며느 receiver부분에 `errorString`에 접근하여, 필드값을 반환합니다.

```go
type errorString struct{
    s string
}
```

- errorString의 구조체 형식

```go
func New(text string) error {
	return &errorString(text)
}
```

- 구조체 `errorString` error패키지의 New()함수를 이용하면 됩니다.
- errors.New("에러값")형태로 호출 시 errorString의 값을 받아서 반환합니다.

```go
package main

import (
	"fmt"
	"errors"
)

func divide(a float32, b float32) (result float32, err error) {
	if b == 0 {
		return 0, errors.New("0으로 나누지마")
	}
	result = a / b
	return
}

func main() {
	var num1, num2 float32
	fmt.Scanln(&num1, &num2)

	result, err := divide(num1, num2)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
```

```
10 0
0으로 나누지마
panic~~
```

- `errors.New`를 이용하여 잘못된 값을 입력되면, main함수의 `err` 변수에 error값을 넣어줍니다.
- 해당 값이 오류이기때문에 panic에서 `0으로나누지마`가 실행 된 후 실행이 종료됩니다.

---

## 에러 출력 및 처리

- pnic을 이용하여 에러를 처리하였지만, 실제로는 다른 방법으로 처리합니다.
- `log` 패키지를 이용하여 에러함수를 이용하여 에러를 처리합니다.
  - `log.Fatal` : 에러로그 출력과 함께 종료됩니다.
  - `log.Print` : 에러로그 출력 후 정상적으로 진행됩니다.
  - `log.Panic` : 에러로그 출력 후 패닉이 뜨며 종료됩니다.

```go
//section16-logPackage.go
package main

import (
    "fmt"
    "log"
)

func divide(a float32, b float32) (result float32, err error) {
    if b == 0 {
        return 0, fmt.Errorf("%.2f으로 나누지마", b)
    }
    result = a / b
    return
}

func main() {
    var num1, num2 float32
    fmt.Scanln(&num1, &num2)

    result, err := divide(num1, num2)

    if err != nil {
        log.Print(err)
    }

    fmt.Println(result)
}
```

```
10 0
2020/10/05 12:11:55 0.00으로 나누지마
0
```

- 0값을 나눌경우 앞의 panic과 뒤에서 사용되는 log값을 이용할때 다른것을 확인할 수 있습니다.
- 차이점
  - 앞의 panic에서와 달리, 해당 값일 출력이 되면서 error기록을 출력합니다.
  - panic에서는 result가 출력되지 않았지만, log를 이용하였을때는 함수 모두가 실행되고 종료되는것을 확인할 수 있습니다.

```go
//section16-log.go
package main

import (
	"fmt"
	"log"
)

func errorChek(n int) (string, error) {
	if n == 1 {
		s := "Goorm"
		return s, nil // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러 리턴
}

func main() {
	s, err := errorChek(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1",s) // Hello 1

	s, err = errorChek(2) // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Print(err)
	}
	fmt.Println("2",s)

	defer func() {
		s, err = errorChek(4) // 매개변수에 4를 넣었으므로 에러 발생
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("defer 4",s)
	}()

	s, err = errorChek(3) // 매개변수에 3을 넣었으므로 에러 발생
	if err != nil {
		log.Panic(err) // defer 함수로 이동
	}
	fmt.Println("3",s)

	// 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
```

```
1 Goorm
2020/10/05 12:17:01 2는 1이 아닙니다.
2
2020/10/05 12:17:01 3는 1이 아닙니다.
2020/10/05 12:17:01 4는 1이 아닙니다.
exit status 1
```

- `errorCheck`함수에서 1 값이 아니면 에러값을 반환하고, 1값일경우 s값을 반환합니다.

1. 매개변수가 1인경우(log.Fatal)
   - 정상적으로 실행되기 때문에 log.Fatal이 실행되지 않고 Goorm값이 정상적으로 출력이 됩니다.
   - `log.Fatal` : 출력 x
   - `fmt.Println("1",s)` : 1 Goorm
2. 매개변수가 2인경우(log.Print)
   - 값이 정상적이지 않아, 오류와 빈값이 반환되어집니다.
   - `log.Print` : 2020/10/05 12:17:01 2는 1이 아닙니다.
   - `fmt.Println("2",s)` : 2
3. 매개변수가 3인경우(log.Panic)
   - `log.Panic` : 2020/10/05 12:17:01 3는 1이 아닙니다. 와 함께 함수 종료
   - `fmt.Println(s)` : 출력x
4. 매개변수가 4인경우(defer - log.Fatal)
   - `log.Fatal` : 2020/10/05 12:17:01 4는 1이 아닙니다.
   - `fmt.Print("4",s)`: 출력 x

```go
package main

import (
    "fmt"
    "log"
)

func errorChek(score int) (int, error) {
    if score >= 0 {
        return score, nil
    }
    return 0, fmt.Errorf("시험 점수를 양의 정수로 입력하세요.")
}

func main() {
    var score int
    fmt.Scanln(&score)

    s, err := errorChek(score)

    if err != nil {
        log.Panic(err)
    }
    fmt.Printf("시험 점수는 %d점입니다.",s)
}
```

간단한 예시로 0 이하의 값을 입력하지 않았을때만 값을 출력하는 예제입니다.
