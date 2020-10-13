# 반복문 - for

---

## for문

다른언어에서 while(조건식) 조건식이 참일때 까지 반복하는 반복문이 존재하지만 Go에서는 while이 존재하지 않습니다<br>
따라서 반복문을 사용하려면 for문을 이용해야 합니다.

```go
for 초기식;조건식;조건변화식{
    반복수행문
}
```

형식으로 이용해야합니다.
<br><br>

간단하게 1~10까지의 합을 구하는 for문을 구현입니다.
`go run section6-1to10.go`

```go
package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Print(sum)
}
```

```
55
```

---

## 조건식만 있는 for문

- while문 처럼 사용 할 수 있습니다.
  `go run section6-while.go`

```go
package main

import "fmt"

func main() {
	num := 2

	for num < 100 {
		fmt.Println(num)
		num *= 2
	}

}
```

```
2
4
8
16
32
64
```

---

## 무한루프

- while(1)를 이용하여 일반적으로 무한루프를 사용하지만, go는 for문에 초기식,조건식,조건변화식을 작성하지 않고 사용할 수 있습니다.
- 테스트에서는 interrupt를 주어 취소할 수 있습니다.
  `go run section6-loop.go`

```go
package main

import "fmt"

func main() {
	for {
		fmt.Printf("무한루프입니다.\n")
	}
}
```

```
무한루프입니다.
무한루프입니다.
.
.
.
무한루프입니다.
무한�^Csignal: interrupt
```

---

## for range문

- for range는 다른 언어의 forEach와 유사합니다.
- 컬렉션으로 부터 하나씩 값을 가져와 차례대로 실행합니다.
- `for 인덱스,요소값 := range 컬렉션이름`을 이용하여 컬렉션의 길이만큼 반복하여 사용합니다.
- 인덱스와 요소값을 인덱스생략('\_,요소값'),요소값생략('인덱스')하여 사용할 수 있습니다.
- 컬렉션의 맵을 이용하면 정수가 아니더라도 다양한 형태로 사용할 수 있습니다.

```go
package main

import "fmt"

func main() {
    var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
    var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
    }
    for _, num := range arr {
		fmt.Printf("값은 %d입니다.\n", num)
    }
    for index := range arr {
		fmt.Printf("arr[%d]입니다.\n", index)
    }
    for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}
}
```

```
arr[0]의 값은 1입니다.
arr[1]의 값은 2입니다.
arr[2]의 값은 3입니다.
arr[3]의 값은 4입니다.
arr[4]의 값은 5입니다.
arr[5]의 값은 6입니다.
값은 1입니다.
값은 2입니다.
값은 3입니다.
값은 4입니다.
값은 5입니다.
값은 6입니다.
arr[0]입니다.
arr[1]입니다.
arr[2]입니다.
arr[3]입니다.
arr[4]입니다.
arr[5]입니다.
banana의 색깔은 yellow입니다.
grape의 색깔은 purple입니다.
apple의 색깔은 red입니다.
```

---

## 실습

### 실습1. 구구단

단 수를 입력받을 int형 변수 dan을 선언합니다.<br>
사용자로부터 dan을 입력받습니다.<br>
7 X 3 = 21 형태로 출력합니다.<br>
1부터 9까지 곱셈을 출력합니다.<br>
`go run section6-training1.go`

```go
package main

import "fmt"

func main() {
	var dan int
	fmt.Scan(&dan)

	for i:=1;i<=9;i++ {
		fmt.Printf("%d X %d = %d\n",dan,i,dan*i)
	}
}
```

```
4 X 1 = 4
4 X 2 = 8
4 X 3 = 12
4 X 4 = 16
4 X 5 = 20
4 X 6 = 24
4 X 7 = 28
4 X 8 = 32
4 X 9 = 36
```

### 실습2. 빛나는 이등변 삼각형

사용자로부터 이등변삼각형의 빗변을 제외한 같은 값의 두 변의 길이를 입력받습니다.<br>
빗면이 \* 모양으로 빛나는 이등변 삼각형이 출력됩니다.<br>
기호와 기호 사이는 띄어쓰기를 합니다.<br>
`go run section6-training2.go`

```go
package main

import "fmt"

func main() {
	var num int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("o ")
		}
		fmt.Println("* ")
	}
}

```

```
*
o *
o o *
o o o *
o o o o *
```
