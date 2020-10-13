# switch

---

## switch문에 의한 선택적 실행

조건에 따른 흐름의 분기가 나누어 집니다.<br>
if~else(if)~else에서도 동일하게 분기를 나눌 수 있습니다. 1,2,3,4 로 실행할때 나머지 값으로 분기를 나누어 사용할 수 있습니다<br>
swich는 해당 값이 1,2,3,4에 따라서 라벨로 나누어 줄 수 있어서 보다 직관적일 수 있습니다<br>

switch는 기본적으로 변수를 가져와 태그로 사용합니다.<br>
변수는 어느 자료형이든 쓸 수 있습니다 태그값이 라벨과 일치하는 값을 찾고 일치하는 case문을 수행합니다.<br>
그리고 case옆에도 라벨링만 하는것이 아니라 조건문도 사용할 수 있습니다.<br>
default는 어떤 값에도 해당되지 않는 값을 수행합니다.<br>
if 처럼 실행문을 중괄호안에 사용하지 않아도 되고, break문으로 라벨링마다 선언해 줄 필요가 없습니다.

```go
switch 태그/표현식{
    case 라벨/표현식:
    실행문
    case 라벨/표현식:
    실행문
    case 라벨/표현식:
    실행문
    default :
    실행문
}

```

```go
package main

import "fmt"

func main() {
	var num int
	fmt.Print("정수 입력:")
	fmt.Scanln(&num)

	switch num {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	default:
		fmt.Println("모르겠어요.")
	}
}
```

C언어에서는 switch를 사용할때 num의 인자는 반드시 필요하고, 정수형 변수여야 합니다.<br>
case옆에 쓰이는 라벨도 상수 형태로만 쓰일 수 있습니다.<br>
즉, 참/거짓을 판별하는 조건식은 쓸 수 없습니다

---

### Go언어 에서의 switch문

- switch에 전달되는 인자로 태그 활용
  - `go run section8-defaultReturn.go`
- switch에 전달되는 인자로 표현식 사용
  - `go run section8-default.go`
- switch에 전달되는 인자 없이 case에 표현식 사용
  - `go run section8-switch,if.go`

```go
//switch에 전달되는 인자로 태그 사용
package main

import "fmt"

func main() {
	var fruit string

	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)

	if (fruit != "apple") && (fruit != "banana") && (fruit != "grape") {
		fmt.Println("잘못 입력했습니다.")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("RED")
	case "banana":
		fmt.Println("YELLOW")
	case "grape":
		fmt.Println("PURPLE")
	}
}
```

```
apple, banana, grape중에 하나를 입력하시오:apples
잘못 입력했습니다.
```

default문을 생략하고 case에 없는값을 입력하여도 에러가 없습니다<br>

- default문을 생략이 가능하지만 if문을 사용하여 조건식을 작성해주어야 하므로, 코드가 길어집니다
- main() 내부에서 return을 실행하면 함수가 종료되는것을 확인할 수 있습니다.

```go
//switch에 전달되는 인자로 태그 사용
package main

import "fmt"

func main() {
    var fruit,color string

	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)

	switch fruit {
	case "apple":
		color="red"
	case "banana":
		color="yello"
	case "grape":
        color="purple"
    default:
        fmt.Println("모르겠습니다")
        return
    }
    fmt.Println(color)
}
```

```
apple, banana, grape중에 하나를 입력하시오: apples
모르겠습니다
```

default문에 return값을 주어서 값을 입력하지 않으면 `모르겠습니다`를 출력한 뒤 하단의 `fmt.Println(color)`부분이 샐행되지 않고 종료 되는것을 알 수 있습니다.

```go
//switch에 전달되는 인자 없이 case에 표현식 사용(참/거짓 판별)
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("정수 a와 b를 입력하시오:")
	fmt.Scanln(&a, &b)

	switch {
	case a > b:
		fmt.Println("a가 b보다 큽니다.")
	case a < b:
		fmt.Println("a가 b보다 작습니다.")
	case a == b:
		fmt.Println("a와 b가 같습니다.")
	default:
		fmt.Println("모르겠어요.")
	}
}
```

```
정수 a와 b를 입력하시오:19 29
a가 b보다 작습니다.
```

switch case문에서 조건식을 사용할 수 있습니다.<br>

---

## 실습

### 실습1 안좋은 계산기

사용자에게 입력받을 연산 번호를 저장할 정수형 변수 sel을 선언합니다.<br>
연산을 할 num1과 num2, 결괏값을 저장할 result를 실수형으로 선언합니다.<br>
1번은 덧셈, 2번은 뺄셈, 3번은 곱셈, 4번은 나눗셈을 연산합니다.<br>
이외의 숫자를 입력할 시에 "잘못된입력입니다."라고 출력하고 프로그램이 종료됩니다.<br>
`go run section8-training1.go`

```go
package main

import "fmt"

func main() {
	var sel int
	var num1, num2, result float32

	fmt.Scanf("%d", &sel)
	fmt.Scanf("%f %f", &num1, &num2)

	switch sel {
	case 1:
		result = num1 + num2
	case 2:
		result = num1 - num2
	case 3:
		result = num1 * num2
	case 4:
		result = num1 / num2
	default:
		fmt.Println("잘못된입력입니다.")
		return
	}
	fmt.Printf("%.1f\n", result)
}
```
