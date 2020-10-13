# 컬렉션

---

## 배열(Array)

두 개 이상의 언어를 모아 놓은 것을 컬렉션이라 합니다.
컬렉션은 다수의 데이터를 저장,처리하는데 유용합니다.

Go는 고정되어진 길이 내에서 연속적으로 데이터를 저장할 수 있습니다. 즉 Go는 배열의 크기는 정적으로 사용할 수 있습니다.
`var 배열이름[배열크기]자료형` 형태로 선언하며, 자료형을 뒤에 선언하는것이 다른 언어와의 차이점을 가지고 있습니다.
또한 배열의 크기는 자료형을 구성하는 한 요소 입니다. int[3],int[5]는 같은 타입의 언어지만 크기가 달라 다른 타입의 언어가 됩니다.
array의 크기를 벗어나는 값을 입력하면 out of bounds for (index)-element array로 문법에러가 발생합니다.

`section10-array.go`

```go
//section-array.go
package main

import "fmt"

func main() {
	var arr1 [5]int   //길이가 5인 int형 배열 arr1을 선언
	fmt.Println(arr1) //숫자를 선언하지 않고 출력해보기

	arr1 = [5]int{1, 2, 3, 4, 5}        //배열 초기화
	fmt.Println(arr1, arr1[0], arr1[4]) //배열 전체와 인덱스에 저장된 값들 출력해보기

	arr2 := [4]int{4, 5, 6, 7} //:= 를 이용해 선언
	arr2[0] = 32               //인덱스를 이용해 값을 초기화
	fmt.Println(arr2)          //arr2 전체 출력해보기

	var arr3 = [...]int{9, 8, 7, 6} //[...]을 이용한 배열 크기 자동 설정
	fmt.Println(arr3, len(arr3))    //arr3 전체와  배열 크기 출력해보기
}
```

Go또한 다른 언어처림 배열의 길이(len),배열의 값변환 `name[index]=data`을 사용할 수 있습니다.
초기화되어진 수 만큼 사용하는 [...]를 이용할 수 있지만 이때는 초기화에서만 동적으로 생성되고, 이후에 값을 추가하는것을 불가능합니다.

### 다차원 배열

2중,3중배열 2차원배열,3차원배열 다른 언어처럼 go도 사용할 수 있으며, 선언 부분에서 크기를 지정한 후 사용할 수 있습니다.

```go
//section10-matrix.go
package main

import "fmt"

func main() {
	var multiArray [2][3][4]int //3차원 배열 선언
	multiArray[1][1][2] = 10    // 인덱스를 이용한 값 초기화package main
    fmt.Println(multiArray)

      var a = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}, //3x3배열 초기화
    }

    fmt.Println(a[1][2]) //2행 3열의 값 출력
}
```

```
[
    [
        [0 0 0 0]
        [0 0 0 0]
        [0 0 0 0]
    ]
    [
        [0 0 0 0]
        [0 0 10 0]
        [0 0 0 0]
    ]
]
6
```

---

## slice

Go는 다른 언어처럼 동적으로 변화,부분적으로 발췌하는 Method지원하지 않습니다.<br>
그러므로, slice를 이용하여 크기를 미리 지정하지 않고 필요에따라서 변경,부분발췌를 기능을 이용합니다.<br>

슬라이스는 다른 자료형의 선언과 다른형태로 선언되어집니다.<br>

```go
// 선언 변수이름 자료형
var num1 int
// 선언 변수이름 자료형[배열크기]
var num2 int[3]

//slice 선언
var num3 []int
```

위의 num1,num2의 선언시 초기화를 따로 명시하지 않았지만 `num1=0`,`num2=[0,0,0]`임을 알 수 있습니다.<br>
허나 slice로 선언되어진 num3는 `nil` 즉 null값을 가집니다.<br>
메모리에 선언과 함께 값을 할당하지만, 슬라이스는 값을 선언하면 배열의 일부분을 기리키는 포인터를 만듭니다.<br>
즉, 선언만 하고 초기화 하지 않으며 슬라이스의 정보만 가지고 있는 배열만 생성됩니다.<br>
<br>
슬라이스는 크기를 지정하지 않기 때문에, 컴퓨터에서는 값을 초기화 하지 못한 상태여서 nil로 출력이 되면 nil과 값이 같은지 검사를 하면 true값을 반환합니다.<br>
`[ptr,len,cap]`로 slice는 이루어져 있으며,

- ptr: 배열의 위치
- len: 배열의 길이
- cap : 전체크기

형태로 이루어져 있습니다. <br>

```go
//section10-slice.go
package main

import "fmt"

func main() {
	var a []int
	fmt.Println("a의값", a, "a의길이", len(a), "cap값", cap(a))
	a = []int{1, 2, 3}
	fmt.Println("a{1,2,3}의값", a, "a의길이", len(a), "cap값", cap(a))

	var b []int
	if b == nil {
		fmt.Println("b is nil")
	}
}
```

```
a의값 [] a의길이 0 cap값 0
a{1,2,3}의값 [1 2 3] a의길이 3 cap값 3
b is nil
```

append로 값을 하나만 추가하였지만, cap의 값은 2배가 증가한것을 볼 수 있습니다.<br>
cap은 따로 지정하지 않으면 초기에 설정되어진 값으로 증가되기 때문입니다.<br>
append시 별도의 cap값을 선언해주지 않았기 때문에 증가 전의 3의 값에 3이 추가되어 cap값은 6이 되었음을 알 수 있습니다.

## make

make는 슬라이스 선언만 하면서 크기를 미리 지정할 수 있습니다.<br>
slice는 선언만 한 후 생성되는 값에따라서 크기가 동적으로 변경되었지만, make는 선언과 동시에 길이와 전체용량을 설정할 수 있습니다.<br>
`make(타입,길이,용량)`형태로 선언되어집니다. 길이랑 용량의 차이는

- len(Length) : 요소의 개수입니다. 추가,삭제가 될때마다 len값은 요소의 갯수에 맞춰서 값이 변경되어집니다.
- cap(Capacity) : 요소가 동적으로 변경되어지기 때문에 메모리 관리를 위해서 사용되어집니다. 길이에 맞게 선언되어진 메모리 용량에 추가를 하기 위해서입니다.

make,slice모두 동일하게 값을 초기화 한 후에 다시 입력하는형태로 값을 추가하면 이전에 존재하였던 값들 대신에 새로운 값들이 대입됩니다.<br>
값을 추가 하기 위해서 `append`를 이용하여 사용할 수 있습니다.

```go
//section10-append.go
package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 4}
	fmt.Println(a, "a의길이", len(a), "a의길이", cap(a))

	b := []int{5}
	fmt.Println(b, "b의길이", len(b), "b의길이", cap(b))

	c := append(a, 5)
	fmt.Println(c, "c의길이", len(c), "c의길이", cap(c))
}
```

```
[1 2 3 4] a의길이 4 a의길이 4
[5] b의길이 1 b의길이 1
[1 2 3 4 5] c의길이 5 c의길이 8
```

```go
//section10-make.go
package main

import "fmt"

func main() {
	a := make([]int, 2, 3)

	for i := 0; i <= 10; i++ {
		a = append(a, i)
		fmt.Println(i, "값 추가시 길이 : ", len(a), "용량 크기 : ", cap(a))
    }
    fmt.Println(a);
}
```

```
0 값 추가시 길이 :  3 용량 크기 :  3
1 값 추가시 길이 :  4 용량 크기 :  6
2 값 추가시 길이 :  5 용량 크기 :  6
3 값 추가시 길이 :  6 용량 크기 :  6
4 값 추가시 길이 :  7 용량 크기 :  12
5 값 추가시 길이 :  8 용량 크기 :  12
6 값 추가시 길이 :  9 용량 크기 :  12
7 값 추가시 길이 :  10 용량 크기 :  12
8 값 추가시 길이 :  11 용량 크기 :  12
9 값 추가시 길이 :  12 용량 크기 :  12
10 값 추가시 길이 :  13 용량 크기 :  24
[0 0 0 1 2 3 4 5 6 7 8 9 10]
```

make를 이용하여 길이2,용량3의 slice를 생성했을때, 0,1,2로 3개의 값이 0으로 초기화 되고, 이후 추가될때마다 용량크기가 2배 증가함을 알 수 있습니다.

## 추가,병합,복사

### 추가

`append`를 이용하여 slice내부에 값을 추가할 수 있습니다.<br>
용량이 남아 있을 경우 길이에 추가하고, 용량의 초과할경우 현재 용량의 크기만큼 추가하여 용량을 2배로 증가하여 사용합니다.<br>
또한 slcie에 slice를 추가하s여 사용할 수 있습니다. 여기서 추가할 slice뒤에는 ...를 입력해주어야 합니다.<br>
즉 `append(sliceA,sliceB...)`형태로 사용할 수 있습니다.<br>
여기서 보면 ...은 javascript의 `Destructuring`과 비슷한 형태임을 알 수 있습니다. 그러므로 `sliceB`를 추가하는 것이 아니라, `sliceB`의 요소를 추가하는것을 알 수 있습니다.
`go run section10-append.go`,`go run section10-appendSlice.go`

```go
package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4}
	sliceB := []int{5, 6, 7, 8}

	sliceA = append(sliceA, sliceB...)
	fmt.Println("sliceA", sliceA)
	fmt.Println("sliceA", sliceB)
}
```

```
sliceA [1 2 3 4 5 6 7 8]
sliceA [5 6 7 8]
```

### copy

slice를 복사할 copy를 이용할 수 있습니다.

```go
package main

import "fmt"

func main() {
	sliceA := []int{0, 1, 2}
	sliceB := make([]int, 2, 4)

	copy(sliceB, sliceA) //A를 B에 붙여넣는다
	sliceB[1] = 10

	fmt.Println(sliceA, "의 길이 :", len(sliceA), "용량 : ", cap(sliceA))
	fmt.Println(sliceB, "의 길이 :", len(sliceB), "용량 : ", cap(sliceB))
}
```

```
sliceA :  [0 1 2] 의 길이 : 3 용량 :  3
sliceB :  [0 10] 의 길이 : 2 용량 :  4
```

copy는 길이,크기를 복사하여 사용하는것이 아니라 요소를 복사하여 이용합니다<br>
또한 복사대상의 길이가 아닌 선언한 값만큼 복사를 진행합니다.<br>
sliceA와 똑같은 길이와,용량,요소값을 복사 하려면 `[시작할위치:복사할크기]`를 이용하여야 합니다.

```go
//section10-copy.go
package main

import "fmt"

func main() {
	sliceA := []int{0, 1, 2}
	sliceB := make([]int, 2, 4) //sliceA에 2배 용량인 슬라이스 선언
	sliceC := make([]int, 2, 4) //sliceA에 2배 용량인 슬라이스 선언

	copy(sliceB, sliceA) //A를 B에 붙여넣는다
	sliceB[1] = 10

	fmt.Println("sliceA : ", sliceA, "의 길이 :", len(sliceA), "용량 : ", cap(sliceA))
	fmt.Println("sliceB : ", sliceB, "의 길이 :", len(sliceB), "용량 : ", cap(sliceB))

	sliceC = sliceA[0:len(sliceA)]

	fmt.Println("sliceC : ", sliceC, "의 길이 :", len(sliceC), "용량 : ", cap(sliceC))

	fmt.Println("sliceA 1번 인덱스부터 출력:", sliceA[1:])
	fmt.Println("sliceA 1번 인덱스까지 출력:", sliceA[:2])

}
```

```
sliceA :  [0 1 2] 의 길이 : 3 용량 :  3
sliceB :  [0 10] 의 길이 : 2 용량 :  4
sliceC :  [0 1 2] 의 길이 : 3 용량 :  3
sliceA 1번 인덱스부터 출력: [1 2]
sliceA 1번 인덱스까지 출력: [0 1]
```

---

## Map

배열,슬라이스는 index즉 값의 요소의 위치를 0,1,2,3,4 형식으로 순서대로 위치를 조회할 수 있습니다.<br>
map는 hash table입니다. 즉 `<키>:<밸류>`형태로 이루어져 있습니다<br>
슬라이스와 맵은 직접적으로 저장하는것이 아닌 참조타입으로 저장합니다<br>
`var 맵 이름 map[key자료형]value자료형`형태로 선언하며, slice처럼 초기화값이 없으면 Nil map형태 입니다.

```go
//section10-map.go
package main

import "fmt"

func main() {
	var a map[int]string

	if a == nil {
		fmt.Println("nil map")
	}

	var m = map[string]string{ //key:value, 형식으로 초기화한다
		"apple":  "red",
		"grape":  "purple",
		"banana": "yellow",
	}

	fmt.Println(m, "\nm의 길이는", len(m))
}
```

```
nil map
map[apple:red banana:yellow grape:purple]
m의 길이는 3
```

## 추가,갱신,삭제

### 추가

slice는 `append`를 이용하여 추가를 하였지만 `맵이름[key]=value`형식으로 값을 추가할 수 있습니다<br>
선언되어진 key값을 이용하면 value값이 변경됩니다.
delete를 이용하여 key값을 삭제할 수도 있습니다.`delete(맵이름,key)`를 이용하면 key,value값이 삭제 됩니다

```go
//section10-mapAddDelete.go
package main

import "fmt"

func main() {
	//지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "충청남도"
	m["053"] = "대구광역시"

	fmt.Println(m)

	//동일한 key값으로 value값을 저장하면 갱신이 된다
	m["032"] = "인천"

	fmt.Println(m)

	//m에 있는 "031"key의 value와 함께 삭제
	delete(m, "031")

	fmt.Println(m)
}
```

```
map[02:서울특별시 031:경기도 032:충청남도 053:대구광역시]
map[02:서울특별시 031:경기도 032:인천 053:대구광역시]
map[02:서울특별시 032:인천 053:대구광역시]
```

## Map의 key체크과 value읽기

`맵이름[key]`를 이용하여 key의 존재여부를 알 수 있으며, 존재할 경우 value값을 반환합니다.

```go
//section10-mapKey.go
package main

import "fmt"

func main() {
	//지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "충청남도"
	m["053"] = "대구광역시"

	val, exist := m["055"]
	fmt.Println("055 : ", val, exist)
	val, exist = m["02"]
	fmt.Println("02 : ", val, exist)
}
```

---

## 실습

### 실습1 역행렬

행렬 A는 2x2 정방 행렬로서 1행에 {7, 3}, 2행에 {5, 2} 원소로 이루어져 있습니다. 이를 배열로 초기화합니다.<br>
변수 'd'에 행렬A의 역행렬 판별식을 초기화합니다. (판별식 = a*d - b*c) <br>
판별식(d)이 0이 아니면(역행렬이 존재할 때) 행렬 A를 역행렬로 만드는 연산을 수행합니다. 그리고 true와 연산 결과인 역행렬을 출력합니다.<br>
판별식(d)이 0이면(역행렬이 존재하지 않을 때) 아무 연산도 하지 않고 false만 출력합니다.<br>

```go
//section10-training1.go
package main

import "fmt"

func main() {
	var a = [2][2]int{
		{7, 3},
		{5, 2},
	}
	var b [2][2]int
	var c bool
	d := a[0][0]*a[1][1] - a[0][1]*a[1][0]
	if d != 0 {
		b[0][0] = 1 / d * a[1][1]
		b[0][1] = 1 / d * -a[0][1]
		b[1][0] = 1 / d * -a[1][0]
		b[1][1] = 1 / d * a[0][0]
		c = d != 0
		fmt.Println(c)
		fmt.Println(b[0][0], b[0][1])
		fmt.Println(b[1][0], b[1][1])
	} else {
		fmt.Println(false)
	}
}
```

```
true
-2 3
5 -7
```

### 실습2 가장 긴 이름

슬라이스 변수 names를 make 함수를 이용해 선언합니다.<br>
사용자에게 입력받는 이름 변수 name을 string형으로 선언합니다.<br>
이름은 엔터(개행)로 제한 없이 입력 받고 1을 입력하면 입력을 종료합니다.<br>
같은 길이의 이름이면 가장 먼저 입력한 이름이 출력됩니다.<br>
가장 긴 이름과 그 길이가 결괏값으로 출력됩니다.<br>

```go
//section10-training2.go
package main

import "fmt"

func main() {
	names := make([]string, 1, 4)

	var name string

	for true {
		fmt.Scanln(&name)
		if name == "1" {
			break
		} else if len(names[0]) < len(name) {
			names[0] = name

		}
	}

	var result string = names[0]

	fmt.Println(result, len(result))
}
```

### 실습3 중간고사 평균 점수

중간고사 과목과 점수를 저장하는 map 컬렉션을 선언합니다.<br>
과목은 string, 점수는 int형으로 선언합니다.<br>
평균값을 저장하는 변수 avg를 float32형으로 선언합니다.<br>
과목과 점수를 제한없이 입력 받습니다. 그리고 0이 입력되면 입력을 종료합니다.<br>
점수들의 평균을 소수점 두 번째 자리까지 출력합니다.<br>

```go
//section10-training3.go
package main

import "fmt"

func main() {
	var m = make(map[string]int)
	var sub string
	var jum int
	var avg float32

	for true {
		fmt.Scanf("%s %d", &sub, &jum)
		if sub == "0" {
			break
		} else {
			m[sub] = jum
			avg += float32(jum)
		}
	}
	for k, v := range m {
		fmt.Printf("%s %d\n", k, v)
	}
	fmt.Printf("%0.2f",avg/float32(len(m)))
}
```

```
Math 73
Science 89
Korean 50
English 98
77.50
```
