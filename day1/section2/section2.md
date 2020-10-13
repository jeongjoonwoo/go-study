# Section2

---

## Constraint and Variable

### 1. 콘솔 출력함수

Go언어 패키지인 fmt를 import하지 않아도 console 출력 해주는 println과 print를 지원합니다.<br>
Print와 Println의 차이점은 개행을 하느냐,안하느냐 차이입니다.<br>
Print는 개행을 하고 Println은 개행을 하지 않습니다<br>

- #### print

  - print는 개행을 하지 않습니다.
  - 개행을 위하여 \n을 입력하여야 사용할 수있습니다.

- #### println

  - println은 실행마다 개행을 진행합니다.<br>
    실습
    `go run section2-print.go`

- #### fmt
  - fmt패키지를 이용하여 콘솔 출력을 할 수 있습니다.
  - `import "fmt"`를 이용하여 fmt패키지를 사용할 수 있습니다.
  - fmt.Print()와,fmt.Println(),fmt.Printf()총 3가지가 존재합니다.
  - Print,Println은 위와 동일하며, Printf는 포맷으로 데이터를 채워서 출력하고자 할때 사용합니다.<br>
    실습
    `go run section2-fmt.go`

---

## Variable Decalre and initialization

### 1. 변수의 선언과 초기화

- Variable

  - 변수를 입력,출력 하고 연산을 위해서 메모리 공간이 필요합니다.
  - Go의 변수 선언방식은 `var <변수이름> <변수형>`입니다.

  ```go
  var var1 int = 1
  var var2 string = "hello"
  var var3 = false
  ```

  - Short Assignment Statement : `:=` 형 선언없이 타입을 추론하여 선언할 수 있습니다.
    - func 즉 ,함수 내부에서만 사용합니다. 전역 변수롤 사용할 함수는 var로 선언 해주어야 합니다.
    ```go
    var var1 := 1
    var var2 := "hello"
    var var3 := false
    ```
  - Go는 변수를 초기화를 하지 않으면 zero value로 설정됩니다.
    - bool타입은 false.숫자는 0 , string ""으로 설정됩니다.
  - Go에서는 선언하고 사용하지 않으면 컴파일 에러가 발생합니다.
    - 패키지,변수,함수 등 모든 선언에서 작용됩니다.
  - 여러개의 변수를 선언하고 한번에 초기화도 가능합니다.

  ```go
  var var1, var2 int = 10, 20
    fmt.Println(var1, var2) //10,20
  ```

  - 다른 언어와의 차이점은 변수형태를 선언 후 변수이름을 선언하지만 Go에서는 변수이름을 선언후 변수형을 선언할 수 있습니다.
  - 실습 `go run section2-var.go`

- Go의 주석
  - go는 주석은 `//`,`/* */`를 이용합니다.
  - `//` : 한줄 주석
  - `/* */` : 단락 주석

### 2. 상수의 선언과 초기화

- Constraint
  - 상수는 초기화 하면 수정이 불가능 합니다.
  - Go의 상수 선언방식은 `const <상수이름> <상수형>`입니다.
  ```go
  const con1 = 1
  ```
  - 상수는 선언시 초기화를 하지 않으면 에러가 발생합니다.
  - 변수와 달리 선언후 사용이 없어도 에러가 출력되지 않습니다.
  - var키워드 대신에 const키워드를 사용하고 생략할수 없기 때문에 `:=`를 사용할 수 없습니다.
  - const는 괄호를 이용하여 여러개를 상수를 한번에 선언할 수 있습니다.
    - 괄호로 같이 묶여있는 상수는 다른형으로 초기화가 가능합니다.
    - 괄호 시작과 마지막의 위치는 상관없지만 상수들은 개행하여 초기화 하여야 합니다.
    - 상수들 사이에 콤마를 사용할 수 없습니다.
    - 묶어서 선언될 경우 첫번째 상수는 꼭 선언되어야 합니다.
    - iota라는 식별자를 값으로초기화하면, 이후 초기화 하지 않고 이어지는 값들은 index로 값이 저장됩니다.
  ```go
  const (
    con1 = 값1
    con2 = 값2
    con3 = 값3
  )
  ```
  실습
  `go run section2-const.go`

---

## 실습

### 실습1. 간단한 덧셈

변수를 메모리에 저장하는 연습을 해보겠습니다. 간단하게 두 변수를 선언하고, 저장한 두 변수에 더하기 연산한 결과 값을 새로운 변수에 초기화해 결과 값을 출력합니다.

```go
package main

import "fmt"

func main(){
	var num1 int =3
	num2 := 7
	sum := num1+num2

	fmt.Printf("%d과 %d의 합은 %d입니다.",num1,num2,sum)
}
```

`go run section2-training1.go`

### 실습2. 잘못된 신상정보

const 상수를 묶어서 선언하는 것을 실습해보겠습니다. const 상수는 프로그램 내에서 특별합니다. 한번 선언되면 수정할 수 없기 때문에 주로 고유값을 초기화 할 때 사용합니다.

그래서 신상정보를 선언하는 코드를 만들어보겠습니다. 출력 결과가 이상할 수 있습니다. 그게 정답입니다. 왜인지 잘 생각해보세요.

```go
package main

import "fmt"

const(
	name = "kim"
	RRN = "800101-1000000"
	job
)

func main() {
	fmt.Printf("%s %s %s",name,RRN,job)
}
```

`go run section2-training2.go`
