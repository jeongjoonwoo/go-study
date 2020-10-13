# 제어문

---

## break,continue,goto

제어문은 주로 반복문에서 쓰입니다. 해당 값에 따라서 분기를 나누기 위해서입니다.<br>
예를들어 1~10까지 출력하다가 3의 배수는 출력하지 않을경우 if문을 사용할 수 있지만 goto문을 이용하여 처리할 수 있습니다.<br>

- break : 반복문을 탈출
- continue : 반복문 첫부분으로 이동
- goto : 특정 부분으로 이동

---

### break

break문이라면 해당 반복문을 탈출하는 의미로 사용할 수 있습니다.
for문을 돌리는 도중에 특정조건이 되면 break로 반복문을 종료할때 사용할 수 있습니다.
다른언어에서는 switch를 이용할때 break문의 선언을 하는 경우가 많으며, select문에서도 사용할 수 있습니다.

1. break는 실행을 하는 도중에 해당 반복문에서 나와 다음 문장을 실행합니다.
   `go run section9-break.go`
2. break 레이블 이름을 사용할 경우 for문을 나와 해당 레이블로 이동하고 break문이 바로 빠져나왔던 for문 다음문장을 실행하게 됩니다.
   `go run section9-breakLabel.go`

```go
package main

import "fmt"

func main() {
	var sum = 0
	var i int

	i = 1

	for {
		sum += i
		if sum > 100 {
			break
		}
		i++
	}

	fmt.Println("1에서 ", i, " 까지 더하면 처음으로 100이 넘어요.")
	fmt.Println("합계:", sum)
}
```

```
1에서  14  까지 더하면 처음으로 100이 넘어요.
합계: 105
```

```go
package main

import "fmt"

func main() {
	var i int = 10

TEST1:
	for {
		fmt.Println(i)
		i -= 1
		if i == 0 {
			break TEST1
		}
	}

	fmt.Println("End")
}
```

```
10
9
8
7
6
5
4
3
2
1
End
```

`Test1`은 아래 for문을 실행하고, for문은 무한루프로 -값을 계속 출력이 되지만
해당 값이 0이 되면 Label에 해당되는 이름을 종료와 함께 다음 값이 출력 되는것을 알 수 있습니다.

---

### continue

continue는 명시한 조건을 걸러주는 역활을 합니다.
continue는 break랑 다른점이, for문과 연관하여서만 사용할 수 있습니다. continue를 만나면 해당 반복문의 조건 검사 부분으로 다시 이동하기 때문입니다.
`go run scetion9-continue.go`

```go
package main

import "fmt"

func main() {
	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			fmt.Printf("%d", i)
		}

		fmt.Printf("%d ", i)
	}
	fmt.Print("\n")
	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
			continue //반복문 처음 부분으로 간다
		}

		fmt.Printf("%d ", i)
	}
}
```

```
1 22 3 44 5 66 7 88 9 1010 11 1212 13 1414 15
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
```

2의 배수의 값들과 아닌값들을 나누어서 출력하는 형태입니다.<br>
continue를 사용하면 1~15까지 한번씩 출력되지만 사용하지 않은 for문은 2의 배수는 2번이 실행 되는것을 확인할 수 있습니다.<br>
즉, continue를 기준으로 아래의 실행문을 실행되지 않는것을 알 수 있습니다.<br>

---

### goto

goto는 흐름을 원하는 위치로 변경하는 키워드입니다.<br>
위치는 label로만 표시를 하는데 `goto label`를 입력하면 해당 label로 이동 되어 사용됩니다.<br>
허나, goto는 프로그램이 흐름을 제어하는 제어문이기 때문에 프로그램의 흐름이 꼬이거나 이상해질 수 있기때문에 사용을 별로 권장하지 않습니다<br>
`go run section9-goto.go`

---

## 실습

### 실습1 구구단2

2단부터 9단까지 출력하는 구구단입니다. 하지만 홀수 단만 출력합니다.<br>
"%d x %d = %d\n"형태로 출력합니다.<br>
단과 단 사이는 한 줄을 비웁니다.<br>
n단은 n x n 까지 출력합니다. 예를 들어, 7단은 7 x 7 = 49까지 출력합니다.<br>
`go run section9-training.go`

```go
package main

import "fmt"

func main() {

	for num1 := 2; num1 < 10; num1++ {
		if num1%2 == 0 {
			continue
		}
		for dan := 1; dan < 10; dan++ {
			if dan > num1 {
				continue
			}
			fmt.Printf("%d x %d = %d\n", num1, dan, num1*dan)
		}
		fmt.Printf("\n")
	}
}
```

```
3 x 1 = 3
3 x 2 = 6
3 x 3 = 9

5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25

7 x 1 = 7
7 x 2 = 14
7 x 3 = 21
7 x 4 = 28
7 x 5 = 35
7 x 6 = 42
7 x 7 = 49

9 x 1 = 9
9 x 2 = 18
9 x 3 = 27
9 x 4 = 36
9 x 5 = 45
9 x 6 = 54
9 x 7 = 63
9 x 8 = 72
9 x 9 = 81

```

### 실습2 두 수를 더하면 99

A와 B는 다른 숫자입니다. 따라서 33 + 66 = 99는 잘못된 출력입니다.<br>
가능한 모든 경우의 수를 출력합니다.<br>
이중 반복문을 사용합니다.<br>
1의 자리에 0을 출력할 수 있습니다.<br>

```go
package main

import "fmt"

func main() {
	var result int

	for num1 := 0; num1 < 10; num1++ {
		for num2 := 0; num2 < 10; num2++ {
			result = num1*10 + num2 + num2*10 + num1
			if result == 99 {
				fmt.Printf("%d%d + %d%d = %d\n", num1, num2, num2, num1, result)
			}
		}
	}
}
```

```
09 + 90 = 99
18 + 81 = 99
27 + 72 = 99
36 + 63 = 99
45 + 54 = 99
54 + 45 = 99
63 + 36 = 99
72 + 27 = 99
81 + 18 = 99
90 + 09 = 99
```
