# 클로저

---

## 외부 변수 접근 : 클로저

`클로저 : 함수안에서 익명함수를 정의하여 바깥함수에 선언되어진 변수에 접근할 수 있는 함수.`

함수 바깥에서 접근하기 위하여 call by value,call by reference를 이용해야 합니다.

```go
//section12-1.go
package main

import "fmt"

func main() {
	a, b := 10, 20
	str := "Hello goorm!"

    result := func() int { // 익명함수 변수에 초기화
        b := 40
        a += b
		return a // main 함수 변수 바로 접근
	}()

	func() {
		fmt.Println(str) // main 함수 변수 바로 접근
	}()

	fmt.Println(result)
	fmt.Println(a,b)
}
```

```
Hello goorm!
50
50 20
```

익명변수에서 a,b,str 를 매개변수로 받지 않았지만 정상으로 50,Hello goorm!이 출력되었습니다.
익명함수가 배개변수를 받지 않았지만 main의 a,b,str을 받아 사용이 되었으며, a값은 수식이 되어져도 값이 변경되었습니다.
익명함수는 변수를 상단에 존재하는 변수를 받아서 사용되어집니다.

```go
package main

import "fmt"

func next() func() int {
	i := 0
	return func() int {
        i += 1
        fmt.Println(&i)
		return i
	}
}


func main() {
	nextInt := next()
    fmt.Println(next())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := next()
    fmt.Println(newInt())
    fmt.Println(nextInt())
}
```

```
0x10a6dc0
0xc0000b4008
1
0xc0000b4008
2
0xc0000b4008
3
0xc0000b4018
1
0xc0000b4008
4
```

1. next함수는 return값으로 func() int를 이용하여 int값을 반환합니다.
2. next함수는 지역변수를 0을 초기화 한 i를 1씩 증가하는 익명함수를 반환합니다.
3. nextInt라는 변수에 next()를 초기화합니다.
   - nextInt는 계속 출력하면 1이 출력되지 않고 2,3으로 증가하는것을 볼 수 있습니다.

i는 익명함수에서 선언되어져서 사용되는것이 아니라 함수 밖의 i를 찾아서 사용됩니다. 또한 nextInt,nweInt는 각각 i값의 주소는 다른위치에 존재합니다.

---

## 실습

### 실습 1. 동전정리

```go
package main

import "fmt"

func calCoin(money int) func(count int) int {
	return func(count int) int {
		return count * money
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)

	if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0 {
		fmt.Println("잘못된입력입니다.")
		return
	}

	add10 := calCoin(10)
	add50 := calCoin(50)
	add100 := calCoin(100)
	add500 := calCoin(500)

	totalmoney := add10(coin10) + add50(coin50) + add100(coin100) + add500(coin500)

	fmt.Println(totalmoney)
}
```

```
1
4
5
6
```

```
3710
```
