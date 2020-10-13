# if/else

---

## 조건에 따른 실행과 분기흐름

Go는 다른 언어에 비하여 조건문이 엄격할 수 있습니다.<br>

### True/False

Go언어의 조건문은 반드시 Boolean형식으로 진행해야 합니다. 그러므로 0,1로 조건문을 사용할 수 없습니다.
`go run section7-boolean.go`

```go
package main

import "fmt"

func main() {
	if 1 {
		fmt.Print("Boolean")
	}
}
```

```
# command-line-arguments
./section7-boolean.go:7:2: non-bool 1 (type int) used as if condition
```

```go
package main

import "fmt"

func main() {
	if true {
		fmt.Print("Boolean")
	}
}
```

```
Boolean
```

---

## 조건식 괄호는 생략이 가능

다른언어에서 몇 가지 언어는 ()는 사용하여 조건문을 넣어주었습니다.<br>
Go는 조건문을 ()를 이용하지 않고 사용할 수 있으며, 실제로 Go에서도 괄호를 생략하고 실행하는것을 권장합니다.<br>
일부 툴에서는 괄호를 넣고 사용하면 자동생략 되는 기능을 지원합니다. (Visual Stuido Code 에서도 지원합니다. Extension/Go를 설치했을때)

## 조건문의 중괄호는 필수

파이썬의 경우는 들여쓰기로 조건문 {}를 이용하고, 실행문이 한줄이면 생략하는 경우가 많습니다.<br>
Go에서는 중괄호를 필수로 사용하여야 합니다.
즉, 블럭안에 실행문이 들어가야지 사용할 수 있습니다.

## else문은 반드시 같은줄에

if~else(if)~else문을 사용할때 사람들마다 사용하는 방식이 다른 경우가 많습니다<br>
else(if)를 사용할때 개행하여 첫글자로 사용하는 방식을 이용할 수 없습니다<br>
그러므로 else문 앞에 중괄호를 작성해주어야 사용할 수 있습니다.

```go
package main

import "fmt"

func main() {
	if 4/2 == 1 { //괄호없이
		fmt.Print("first")
	}
	else{
		fmt.Println("False")
	}
}
```

```
# command-line-arguments
./section7-if,else.go:10:2: syntax error: unexpected else, expecting }
```

`go run section7-if,else.go`

```go
package main

import "fmt"

func main() {
	if 4/2 == 1 { //괄호없이
		fmt.Print("first")
	}
	else if 5/2 == 0 {
		fmt.Println("second")
	} else {
		fmt.Println("last")
	}
}
```

```
# command-line-arguments
./section7-if,else.go:10:2: syntax error: unexpected else, expecting }
```

```go
package main

import "fmt"

func main() {
	if 4/2 == 1 { //괄호없이
		fmt.Print("first")
	} else if 5/2 == 0 {
		fmt.Println("second")
	} else {
		fmt.Println("last")
	}
}
```

```
last
```

## 조건식에 간단한 문장 실행 가능(Optional Statement)

Go언어에서는 실행문이 실행전에 간단하게 문장형식으로 조건식을 사용할 수 있습니다.<br>
`if val:=num*2;val==2`와 `if num==1`은 같은 결과값을 출력합니다.<br>
단, 조건식이 실행하기 전에 val은 if문 블락 내부에서만 사용할 수 있습니다.
`go run section7-optionalStatement.go`

```go
package main

import "fmt"

func main() {
	num := 1
	val := 10
	if num == 1 { //괄호없이
		fmt.Println("num==1")
	}

	if val := num * 2; val == 2 { //괄호있이
		fmt.Println("val==2")
		fmt.Println("in if block val = ", val)
	}

	fmt.Println("out if block val = ", val)
}
```

```
num==1
val==2
in if block val =  2
out if block val =  10
```

---

## 실습

### 실습1 7과 9의 배수

if ~else와 관련된 문장들의 이해보다 중요한 것은 적용능력입니다. 실력 향상을 위해 여러 실습 문제를 직접 풀어보세요.

1이상 100미만의 정수 중에서 7의 배수와 9의 배수를 출력하는 프로그램을 작성해봅니다. 단 7의 배수이면서 동시에 9의 배수인 정수는 한 번만 출력해야합니다<br>
`go section7-training1.go`

```go
package main

import "fmt"

func main() {

	for val:=1;val<100;val++{
		if val%7==0 || val%9==0{
			fmt.Printf("%d ",val)
		}
	}
}
```

```
7 9 14 18 21 27 28 35 36 42 45 49 54 56 63 70 72 77 81 84 90 91 98 99
```

### 실습2 두 수의 차

두 개의 정수를 입력 받아서 두 수의 차를 출력하는 프로그램을 구현해봅니다. 이때, 무조건 큰 수에서 작은 수를 뺀 결과를 출력해야합니다. 따라서 출력 결과는 무조건 0 이상이 되어야 합니다.

int형 변수 num1, num2, result를 선언합니다.<br>
num1과 num2는 연산에 사용되고 result는 결괏값을 저장합니다.<br>
`go run section7-training2.go`

```go
package main

import "fmt"

func main() {
	var num1, num2, result int

	fmt.Scanf("%d %d", &num1, &num2)

	if num1 > num2 {
		result = num1 - num2
	} else {
		result = num2 - num1
	}

	fmt.Print(result)
}
```

input

```
19 2
```

output

```
17
```
