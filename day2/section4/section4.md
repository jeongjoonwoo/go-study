# 자료형

---

## data type

변수는 데이터 저장을 위해 할당된 메모리 공간의 이름입니다. <br>
메모리 할당 이전에 바이트 크기를 정하여 저장하기 위해서 어떤 자료형인지 표현하는것이 자료형입니다.
Go언어는 bool,string,int,float,복소수등이 있습니다.<br>
바이트 크기는 import "unsafe"를 선언한 후 `unsafe.sizeof변수`형태로 크기를 알 수 있습니다.

### 1. 불리언 타입

- 참(True)/거짓(False)를 나타내는 자료형입니다.
- 다른언어와 다르게 0,1로 사용할 수 없습니다.
- size: 1byte

### 2. 정수 타입

- un을 붙이면 C언어의 unsigned처럼 사용할 수 있습니다.
- uinpter 를 포인터의 주소를 할당할때 사용합니다.
  - 크기 : 8byte
  - 비트 패턴을 할당할만한 크기의 사이즈 입니다.
- 크기 : 1byte(int8,uint8),2byte(int16,uint16),4byte(int32,uint32),8byte(int64,uint64)
- int와 uint는 시스템의 환경마다 4byte(32bit system),8byte(64bit Ststem) 로 다릅니다.

### 3. 실수,복소수 타입

- 정밀도 : 예를들어 무한 함수일 경우 해당 함수가 일정 크기를 넘어가면 아래부분은 제외하고 저장하기 위한 값입니다.
- 복소수 선언은 3+4i형태로 선언할 수 있습니다.
- 크기 : 4byte(float32),8byte(float64,complex64),16byte(complex128)

### 4. 문자열 타입

- immutable : 문자열타입은 선언 후 값을 수정할 수 없습니다.
- 크기 : 16byte

### 5. 기타 타입

- 기타 타입은 byte,rune이 있습니다.
- byte값을 8bit 부호없는 정수값과 구별하는데 사용됩니다.
- rune는 int32와 같은 자료형으로 볼 수 있습니다.

---

## String

### 문자열을 표현하는 방법

1. 백틱(``)를 이용하면 Raw String Literal이라고 불립니다.
   - 백틱 내부에 기호를 사용하면 문자열로 인식됩니다.
2. 쌍따옴표("")를 이용하면 Interpreterd String Literal이라고 불립니다.
   - 쌍따옴표 내부에 기호를 사용하면 기호로 디해됩니다.

- +연산자는 숫자 뿐만 아니라 문자열도 합할 수 있습니다.
  `go run section4-plusString.go`

```go
package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	// Interpreted String Literal
	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}

```

```
바로 실행해보면서 배우는 \n Golang

바로 실행해보면서 배우는
Golang

구름 EDU
Golang
```

---

## trans type

### 자료형의 변환

자료형을 사용하는 도중에 다른 타입으로 선언할 상황이 있습니다.<br>

일반적으로 자동형변환 (자동으로 형 변환),강제형변환(형 변환을 병시적으로 작성하여 강제로 변환) 가 존재합니다.<br>
Go언어에서는 형변환을 할때 명시적으로 지정해주어야 합니다.<br>
정수타입을 수식하면, 연산의 결과값과 그것을 저장하려는 타입이 다르면 런타임 에러가 발생합니다
`go run section4-type.go`

```go
package main

import "fmt"

func main() {
	var num int = 10
	var changef float32 = float32(num) //int형을 float32형으로 변환
	changei := int8(num)               //int형을 int8형으로 변환

	var str string = "goorm"
	changestr := []byte(str) //바이트 배열
	str2 := string(changestr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)
}
```

```
10
10 10
goorm
[103 111 111 114 109]
goorm
```

```go
package main

import "fmt"

func main() {
	var num1, num2 int = 3, 4

	var result float32 = float32(num1) / float32(num2)

	fmt.Printf("%f", result)
}

```

```
0.750000
```

---

## 실습

### 실습1 강제 형 변환

강제 형 변환을 실습해보는 예제입니다. 실제로 음수를 uint 형으로 강제 형 변환을 하는 것은 드문 일이지만 형 변환의 과정을 이해한다고 생각하고 연습해봅니다.

int 형 변수 세 개를 사용자로부터 입력받습니다. 이때, 두 번째 숫자는 음수입니다. 그리고 각 변수는 float32, uint, int64 형으로 강제 형 변환이 되어 자료형과 함께 값이 출력됩니다.

`3 -10 5`를 순차적으로 입력합니다.

```go
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int

	fmt.Print("")
	fmt.Scanln(&num1, &num2, &num3)

	var num4 = float32(num1)
	var num5 = uint(num2)
	var num6 = int64(num3)

	fmt.Printf("%T, %f\n%T, %d\n%T, %d\n",num4,num4,num5,num5,num6,num6)
}
```

```
float32, 3.000000
uint, 18446744073709551606
int64, 5
```
