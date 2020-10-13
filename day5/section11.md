# 함수

---

## Go언어에서의 함수

- 함수 : 특정기능을 위해 여러 문장을 묶어 실행하는 코드블럭 단위

프로그램에서 기능별로 특징을 묶어 구현해 놓은 것이 함수입니다. `func 함수이름(매개변수이름 매개변수형)반환형`이 함수 선언 형식입니다.<br>

- 함수 선언할때 키워드는 func입니다
- 반환형이 괄호 뒤에 명시됩니다. 매개변수형도 매개변수이름 뒤에 선언됩니다.
- 함수는 패키지안에서 정의되고 호출되는 함수가 꼭 호출하는 함수앞에 있을 필요가 없습니다.
  - Go는 객체지향 언어이므로 선언시 함수가 반드시 앞에 있어야할 필요가 없습니다.
- 반환값이 여러개일경우 `((반환형1,반환형2))`로 선언해야 합니다.
  - 반환값이 1개 이상 사용시 필요합니다.
- 블록 시작 괄호가 함수 선언 동시에 첫줄에 있어야 합니다.
- 호출은 `함수이름(전달인자이름)`형식으로 입력해야합니다.

함수는 3가지 형태로 나눌 수 있습니다.

1. 매개변수,반환값이 있는 형태
2. 매개변수는 있지만 반환값이 없는 형태
3. 매개변수는 없지만 반환값이 있는 형태
4. 매개변수,반환값이 없는 형태

---

## 전역변수와 지역변수

매개변수는 2가지(call by value, call by referenece) 형태가 있습니다. call by reference는 이미 map,slice에서 배운 개념입니다. 주소값을 받아서 사용하는 형태입니다.<br>
<br>
`전역변수`,`지역변수`가 중요한 이유는 매개변수 전달시 형태에 따라서 결과값이 다르기 때문입니다.<br>
<br>
`지역변수`와 `전역변수`는 선언위치에 따라 유형이 결정됩니다.
`지역변수`는 중괄호 안에서 선언된 변수를 의미합니다. 즉, `지역변수`의 영역은 함수내 한정되어 중괄호 외부에서 사용할 수 없습니다.<br>
`전역변수`는 중괄호 외부에 선언된 변수를 의미합니다. `지역변수`와 달리 여러 함수 에서 사용할 수 있습니다.<br>
`지역변수`,`전역변수`는 이외에도 차이점이 있습니다. 메모리에 존재하는 시간이 다릅니다.<br>
`지역변수`는 해당 함수 내에서만 사용이 되니 함수가 실행될때 메모리영역을 차지한 후 함수가 종료되면 소멸됩니다.<br>
`전역변수`는 프로그램이 시작시 메모리를 차지하며, 종료될때까지 유지됩니다.

```go
//section11-local.go
package main

import "fmt"
func exampleFunc1() {
	var a int = 10 //지역변수 선언
	a++
	fmt.Println("exampleFunc1의 a는 ", a)
}

func exampleFunc2() {
	var b int = 20 //지역변수 선언
	var c int = 30 //지역변수 선언
	b++
	c++
	fmt.Println("b와 c는 ", b, c)
}

func main() {
    var a int = 28
	exampleFunc1()
	exampleFunc2()
	fmt.Println("main의 a는", a)
}
```

```
exampleFunc1의 a는  11
b와 c는  21 31
main의 a는 28
```

1. main함수를 실행합니다.
   1. `var a int = 28`를 실행하며 메모리영역에 해당 값을 저장합니다.
2. `exampleFunc1()`를 실행합니다.
   1. 실행하며, `var a int =10`를 메모리에 저장합니다.
   2. 2-1의 a값을 1 증가한 후 저장합니다.
   3. 2-1에 선언되어진 a값이 출력됩니다.
   4. 메모리에 2-1변수는 함수 종료와 함께 삭제됩니다.
3. `exampleFunc2()`를 실행합니다.
   1. `var b int = 20 var c int = 30 `를 메모리에 저장합니다.
   2. b,c를 각각 1증가한 후 저장합니다.
   3. b,c를 각각 출력합니다.
   4. 메모리에 2-1변수는 함수 종료와 함께 삭제됩니다.
4. `fmt.Println("main의 a는",a)`를 실행합니다.
   1. 메모리에 저장되어진 a=28을 출력합니다.

```go
//section11-global.go
package main

import "fmt"

var d int = 100

func exampleFunc1() {
    d++
    fmt.Println("d값",d)
}

func exampleFunc2() {
	var d int = 10
	d++
	fmt.Println("d값",d)
}

func main() {
	exampleFunc1()
	exampleFunc2()
	fmt.Println("main의 d는", d)
}
```

```
d값 101
d값 11
main의 d는 101
```

1. 실행과 동시에 메모리에 `var d int = 100`를 저장합니다.
2. `exampleFunc1`이 실행되면 d=100이므로 이를 1증가후 해당 메모리에 저장합니다.
3. `exampleFunc2`이 실행되면 d=10를 지역내에서 새로 선언하여 다른 주소값에 저장합니다.
   1. d=10을 1 증가후 함수종료와 함께 메모리에서 지워집니다.
4. main에는 `exampleFunc1`에서 실행된 1증가값과 함꼐 d값이 101이 출려되며 종료가 됩니다.

---

## 매개변수

지역변수로 선언되어진 변수를 다른 함수에 전달하여 해당 값을 함수에 전달하기 위하여 사용하는것이 매개변수입니다.
`func 함수이름(매개변수이름 매개변수형)반환형`에서 매개변수입니다.

### Pass by value

```go
//section11-passByValue
package main

import "fmt"

func printSqure(a int) {
	a *= a
	fmt.Println(a)
}
func main() {
	a := 4 //지역변수 선언
	printSqure(a)
	fmt.Println(a)
}
```

```
16
4
```

매개변수를 받아 제곱하는 printSqure에 4를 매개변수로 보내준후 16이 출력되는것을 확인하였으나, main함수에는 4가 출력되었습니다<br>
이는 매개변수가 복사가 된 후 전달되기 때문에 a값이 변환되지 않고 유지되기 때문입니다.<br>
즉, call by value는 값을 복사하여 전달하기 때문에 main의 a와 printSqure의 a는 같은값이지만 다른주소를 가집니다.

---

### pass by reference

```go
package main

import "fmt"

func printSqure(a *int) {
	*a *= *a
	fmt.Println(*a)
}
func main() {
	a := 4 //지역변수 선언
	printSqure(&a) //참조를 위한 a의 주솟값을 매개변수로 전달
	fmt.Println(a)
}
```

```
16
16
```

앞의 value와 똑같은 역활후 코드의 변화가 없지만 전달하는 매개변수의 형태가 다릅니다.<br>
value는 값을 복사하여 전달을 하지만, 아래는 `pointer`를 이용하여 변수의 주소를 전달합니다.<br>
즉, printSquer의 a와 main의 a는 같은 주소의 값을 가진다는 의미입니다. 그러므로 함수에서는 메모리영역을 새로 확보하는 것이 아니라 main의 a를 받아 함수가 실행됩니다.<br>

---

### 가변인자함수

`가변인자함수`는 매개변수의 수가 일정하지 않을때 사용됩니다. 예를들어 값을 더해주는 함수가 있을경우, 2개를 합할 수 도 있고, 10개의 수를 보내여 합할수도 있습니다. 이때 매개변수를 10개 선언할 필요없이 가변인자함수를 이용하여 여러개의 함수를 받아서 사용할 수 있습니다.
허나, 가변인자 함수도 몇가지 제한사항이 있습니다.

1. 같은 형태의 매개변수를 전달해야합니다.
2. 전달되어진 변수들은 슬라이스 형태입니다.

`func 함수이름(매개변수이름 ...매개변수형) 반환형`,`함수이름(슬라이스이름...)`형태로 이용할 수 있습니다.

```go
//section11-variableArgument.go
package main

import "fmt"

func addOne(num ...int) int {
	var result int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
	}
	return result
}

func addTwo(num ...int) int {
	var result int

	for _, val := range num { //for range문을 이용한 num의 value 순차 접근
		result += val
	}
	return result
}

func main() {
	num1, num2, num3, num4, num5 := 1, 2, 3, 4, 5
	nums := []int{10, 20, 30, 40}

	fmt.Println(addOne(num1, num2))
	fmt.Println(addOne(num1, num2, num4))
	fmt.Println(addOne(nums...))
	fmt.Println(addTwo(num3, num4, num5))
	fmt.Println(addTwo(num1, num3, num4, num5))
	fmt.Println(addTwo(nums...))
}
```

```
7
100
12
13
100
```

---

## 반환값

프로그래밍 원본값은 저장한 후 해당 값만 변형되어 다시 사용하기 원할때가 존재합니다<br>
즉 반환값은 pass by value에서 값을 전달한 후, 해당 값을 받기 위해서 사용됩니다.<br>
앞의 [가변인자](###가변인자함수)에서 예시로 `return`형식을 이용하여 사용할 수 있습니다.<br>

- 반환값은 개수 만큼 변환형을 명시해아합니다.
  - n개이상일 경우 `(())`형태를 이용해야 합니다.
- 동일한 형태라도 모두 명시해야합니다.

```go
//section11-return.go
package main

import "fmt"

func add(num ...int) (int, int) {
	var result int
	var count int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
		count++
	}

	return result, count
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println(add(nums...))
}
```

```
150 5
```

### Name Return Parameter

여러개의 값을 반환할 경우 괄호안에 반환형을 명시하지 않고, 반황형과 반환값의 이름이 같은 경우 이용합니다.<br>

- `(반환값이름 반환형,...,반환값이름 반환형)`형태로 입력합니다.
- `반환값이름 반황형`자체가 변수입니다.
- return이 반드시 필요합니다.
- 반환값이 하나라도 반환값 이름에 명시되었다면 괄호안에 써야합니다.

```go
//section11-returnName.go
package main

import "fmt"

func dessertList(fruit ...string) (count int, list []string) { //여기서 이미 선언된 것이다

	for i := 0; i < len(fruit); i++ {
		list = append(list, fruit[i])
		count++
	}

	return //생략하면 안 된다
}

func inputFruit() (list []string) { //Named return parameter는 값이 하나라도 괄호를 써야한다

	for {
		var fruit string
		fmt.Print("과일을 입력하세요:")
		fmt.Scanln(&fruit)

		if fruit != "1" {
			list = append(list, fruit)
		} else {
			fmt.Println("입력을 종료합니다.\n")
			break //반복문을 빠져나간다
		}
	}

	return
}

func main() {
	fmt.Println("디저트로 먹을 과일을 입력하고 출력합니다. \n1을 입력하면 입력을 멈춥니다.\n")
	count, list := dessertList(inputFruit()...) //함수를 변수처럼 사용할 수 있습니다
	fmt.Printf("%d개의 과일을 입력하셨고, 입력한 과일의 리스트는 %s입니다.\n", count, list)
}
```

```
디저트로 먹을 과일을 입력하고 출력합니다.
1을 입력하면 입력을 멈춥니다.

과일을 입력하세요:사과
과일을 입력하세요:1
입력을 종료합니다.

1개의 과일을 입력하셨고, 입력한 과일의 리스트는 [사과]입니다.
```

---

## 익명함수

함수는 시작시 메모리를 차지하고, 또한 기능을 수행할때마다 해당 메모리 영역을 찾아가 호출해야하므로 속도가 낮아집니다.<br>
이를 보완하기 위해 사용되는것이 `익명함수`입니다.

```go
//section11-anonymous.go
package main

import "fmt"

func main() {
	func() { //1
		fmt.Println("hello")
	}()

	func(a int, b int) {//2
		result := a + b
		fmt.Println(result)
	}(1, 3)

	result := func(a string, b string) string {//3
		return a + b
	}("hello", " world!")
	fmt.Println(result)

	i, j := 10.2, 20.4//4
	divide := func(a float64, b float64) float64 {
		return i / j
	}(i, j)
	fmt.Println(divide)
}
```

```
hello
4
hello world!
0.5
```

1. main함수 실행되면 1번이 실행되며 `fmt.Println('hello')`가 실행된후 종료됩니다.
2. `func(a int b int)`가 실행되며, 이후 `(1,3)`의 매개변수를 바로 넣어주어 함수를 실행합니다.
3. result는 `func(a string, b string) string`함수를 가지며, 해당 함수를 실행할 수 있습니다.
4. deviede는 함수에서 i에서 j를 나눈 값을 변수로 가집니다.

```go
//section11-funcVariable.go
package main

import "fmt"

func addDeclared(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

func main() {
	var nums = []int{10, 12, 13, 14, 16}

	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}
	fmt.Println(addAnonymous(nums...))
	fmt.Println(addDeclared(nums...))
}
```

```
65
65
```

함수를 변수에 넣어서 사용할 수 있습니다. 익명함수,선언함수는 모두 같은 수식을 진행하며 결과값도 동일한 값을 반환합니다.

```go
//section11-declare,anonymous.go
package main

import "fmt"

func add() {
	fmt.Println("선언 함수를 호출했습니다.")
}

func main() {

	add := func() {
		fmt.Println("익명 함수를 호출했습니다.")
	}

	add()
}
```

```
익명 함수를 호출했습니다.
```

익명함수와 선언함수는 읽는 순서가 다릅니다. 위의 선언함수는 프로그램 시작과 모두 읽습니다. 허나 익명함수는 그 main에서 실행되기 때문에 함수가 실행되는 곳에서 읽습니다<br>
전역변수와 지역변수의 차이저럼 `선언 함수`가 `익명함수`보다 더 먼저 읽힙니다.

### 일급함수

익명함수는 Go의 함수가 일급함수이기 때문에 가능합니다.
`일급함수 : 함수를 기본 타입과 동일하게 사용할 수 있어, 함수 자체를 다른 함수의 매개변수로 전달하거나 다른 함수의 반환값으로 사용될 수 있습니다.`

```go
//section11-firstClassFunc.go
package main

import "fmt"

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}

	r1 := calc(multi, 10, 20)
	fmt.Println(r1)

	r2 := calc(func(x int, y int) int { return x + y }, 10, 20)
	fmt.Println(r2)
}
```

```
200
30
```

calc는 함수, 정수형 변수 2개를 받습니다.
multi는 익명함수로 정수형 변수 2개를 곱한 후 반환합니다.

- r1
  - 함수인 multi, 10,20을 보내어 두개의 값을 곱하여 반환합니다.
- r2
  - 익명함수인 2개의 매개변수를 받아 합하는 함수를 보냅니다. calc는 보내어진 함수를 받아 10,20을 합한후 반환합니다.

### type문을 사용한 함수 원형 정의

`func calc(f func(int, int) int, a int, b int) int`에서 `f func(int,int)int`는 가독성을 떨어트리고 깔끔하지 못한 느낌을 받습니다<br>
만약, 변수가 2개이 상일 경우 `func calc(f func(int, int,int) int, a int, b int) int`형태로 선언되어 집니다, 이때 함수 형태를 정의하고 간단하게 사용할 수 있게 해주는것이 `type`입니다.

```go
//section11-type.go
package main

import "fmt"

//함수 원형 정의
type calculatorNum func(int, int) int
type calculatorStr func(string, string) string

func calNum(f calculatorNum, a int, b int) int {
	result := f(a, b)
	return result
}

func calStr(f calculatorStr, a string, b string) string {
	sentence := f(a, b)
	return sentence
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}
	duple := func(i string, j string) string {
		return i + j + i + j
	}

	r1 := calNum(multi, 10, 20)
	fmt.Println(r1)

	r2 := calStr(duple, "Hello", " Golang ")
	fmt.Println(r2)
}
```

```
200
Hello Golang,Hello Golang
```

`type`으로 형식을 선언하였습니다.`func calNum(f func(int,int) int, a int, b int) int `로 선언하여 사용되어져야 할 부분에서
`func calNum(f calculatorNum, a int, b int) int`로 보다 간편하게 적혀있습니다<br>
`type calculatorNum func(int, int) int` 2개의 매개변수를 받아 하나의 값을 반환한다는 의미로 사용되어집니다.<br>
즉, 일급함수를 사용할 경우 가독성을 높여주기 위하여 type문을 이용합니다.

---

## 실습

### 실습1 오름차순 정렬

inputNums() 함수는 사용자에게 몇개의 정수를 입력 받을지, 입력받은 정수를 슬라이스에 저장하는 기능을 합니다. 입력이 끝나면 int형 슬라이스를 반환합니다.<br>
bubbleSort() 함수는 사용자에게 입력받은 슬라이스를 오름차순 정렬하는 연산을 수행합니다. 힌트는 temp 변수를 사용합니다.
<br>
output() 함수는 정렬이 끝난 슬라이스를 요소 한개씩 띄어쓰기로 출력하는 기능을 수행합니다.
<br>
main() 함수는 3줄입니다.
<br>

```go
//section11-training1.go
package main

import "fmt"

func bubbleSort(num ...int) {
	var temp int
	for i := 0; i < len(num)-1; i++ { //for문을 이용한 num[i] 순차 접근
		for j := i + 1; j < len(num); j++ {
			temp = num[i]
			if num[j] < temp {
				num[i] = num[j]
				num[j] = temp
			}
		}
	}
	return
}

func inputNums() []int {
	var num int
	var data int
	var nums []int
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		fmt.Scanln(&data)
		nums = append(nums, data)
	}
	return nums
}

func outputName(num ...int) {
	for _, v := range num {
		fmt.Printf("%d ", v)
	}
	return
}

func main() {
	nums := inputNums()
	bubbleSort(nums...)
	outputName(nums...)
}
```

### 실습2 아이패드를 사주는 조건

inputNums() 함수는 사용자에게 시험 점수 5개를 입력 받습니다.<br>
calExam() 함수는 총점, 90점 이상 과목 수, 70점 미만 과목 수를 계산 후 모두 반환합니다.<br>
printResult() 함수는 결괏값을 가지고 아이패드를 살 수 있는지 없는지 여부를 출력하고, 불가능하다면 모든 이유도 함께 출력합니다.<br>
main() 함수는 두 줄입니다. calExam() 함수 자체를 printResult() 함수 전달 인자로 사용합니다.<br>

```go
//section11-training2.go
package main

import "fmt"

func inputNums() []int {
	var num int
	var nums []int

	for i := 0; i < 5; i++ {
		fmt.Scanln(&num)
		nums = append(nums, num)
	}

	return nums
}

func calExam(arr []int) (int, int, int) {
	var over []int
	var low []int
	var result int
	for _, value := range arr {
		if value >= 90 {
			over = append(over, value)
		}
		if value < 70 {
			low = append(low, value)
		}
		result += value
	}
	return result, len(over), len(low)
}

func printResult(sum int, over int, low int) {
	var result bool = true

	if sum < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	}
	if over < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	}
	if low > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}

	if result {
		fmt.Println("아이패드를 살 수 있습니다.")
	} else {
		fmt.Println("아이패드를 살 수 없습니다.")
	}
}

func main() {
	nums := inputNums()
	printResult(calExam(nums))
}
```

### 실습3 역학적 에너

위치 에너지에 사용할 중력가속도(9.8)를 상수로 초기화합니다.<br>
위치 에너지를 계산하는 익명 함수를 kinEnergy 변수에 초기화 하고, 위치 에너지를 계산하는 익명 함수를 potEnergy 변수에 초기화합니다. 두 익명 함수는 main() 함수 내에 선언합니다.<br>
calMechEnergy() 함수에 매개변수로서 함수를 사용하기 위해 type문으로 위 두 함수의 원형을 정의합니다.<br>
calMechEnergy() 함수는 전역으로 선언합니다. 매개변수로는 함수, float32형 변수 두개가 선언되고, 반환형은 float32입니다. 매개변수 함수가 kinEnergy면 운동 에너지를, potEnergy 면 위치 에너지를 반환합니다. <br>
사용자에게 질량, 속도, 높이를 차례로 입력받습니다.<br>
calMechEnergy() 함수의 결괏값을 ke와 pe 변수에 각각 초기화합니다. ke는 운동 에너지, pe는 위치 에너지입니다.<br>
최종적으로 운동 에너지, 위치 에너지, 역학적 에너지를 출력합니다. <br>

```go
//section11-training3.go
package main

import "fmt"

const g float32 = 9.8

type calType func(float32, float32) float32

func calMechEnergy(f calType, m float32, v float32) float32 {
	result := f(m, v)
	return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m float32, v float32) float32 {
		return m * v * v / 2
	}
	potEnergy := func(m float32, h float32) float32 {
		return m * g * h
	}

	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)
	fmt.Println(ke, pe, ke+pe)
}
```
