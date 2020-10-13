# 콘솔 출력과 입력 함수

---

## Print

Go언어에서는 fmt패키지 콘솔 출력 함수는 Print,Println,Printf 3가지가 기본적으로 존재하고
파일 출력을 위한 Fprint,Fprintln,Fprintf와 문자열로 반환되는 Sprint,Sprintln,Sprintf가 있습니다.<br>

### fmt package

- 코드 상단에 `import "fmt"`를 선언하여 fmt패키지를 이용할 수 있습니다.
- printf,scanf를 이용하여 해당 값들을 형식화한 입출력함수를 사용할 수 있습니다.

### Println,Print

1. Print

- 개행을 자동으로 추가하지 않습니다.
- 개행을 하기위해서는 \n를 이용해야합니다.

2. Println

- 개행을 자동으로 추가합니다.
  `go run section5-print.go`

```go
package main

import "fmt"

func main() {
	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(n) //개행을 위해 \n를 입력한다.
	fmt.Println(arr)
	fmt.Println("여기까지 Println입니다.")

	fmt.Print(n) //개행을 위해 \n를 입력한다.
	fmt.Print(arr)
	fmt.Print("여기까지 Print입니다.")
}
```

```
14
[1 2 3 4 5]
여기까지 Println입니다.
14[1 2 3 4 5]여기까지 Print입니다.
```

3. Printf

- 개행을 자동으로 추가하지 않습니다.
- 개발자가 포맷팅을 할 수 있습니다.
- 선언을 할때 반드시 서식문자를 이용하여 출력해주어야 합니다.
- 배열을 출력할때 한번에 출력할 수 있습니다.
  `go run section5-printf.go`

```go
package main

import "fmt"

func main() {
	name, age := "Kim", 24
	fmt.Printf("안녕하세요 반갑습니다.")
	fmt.Printf("안녕하세요 반갑습니다.\n")

	fmt.Printf("안녕하세요 저는 %s 입니다. 나이는 %d살 입니다\n", name, age)

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d 배열 입니다\n", arr)
}
```

```
안녕하세요 반갑습니다.안녕하세요 반갑습니다.
안녕하세요 저는 Kim 입니다. 나이는 24살 입니다
[1 2 3 4 5] 배열 입니다
```

### 서식종류

1. %t : bool
2. %c : character
3. %b : 2진수
4. %o : 8진수
5. %d : 10진수
6. %x : 16진수
7. %f : 실수
8. %s : 문자열
9. %p : 포인터
10. %T : 타입
11. %v : 모든 형식

- 원하는 출력폭을 정할때 `%(폭)d`를 이용하여 지정한만큼 폭을 지정할 수 있습니다.
- `%0(폭)d` : 0을 폭만큼 넣습니다.
- `%-(폭)d` : 왼쪽부터 출력을 할 수 있습니다.
- `%.(자릿수)d`: 소수점이하 자릿수를 지정할 수 있습니다.
  `go run section5-printType.go`

```go
package main

import "fmt"

func main() {
	fmt.Printf("공백 네 칸: %4d\n", 10)
	fmt.Printf("4자리중 10을 제외한 앞 2자리는 0으로 채움: %04d\n", 10)
	fmt.Printf("총 4자리중 각각 왼쪽에 정렬 됩니다.: %-4d%-4d\n", 10, 15)
	fmt.Printf("12.3456를 소수점 둘째 자리까지만 표시하면 %.2f입니다.\n", 12.3456)
}

```

```
네 칸 차지하는 13:   13
빈칸은 0으로 채우고 4칸 차지하는 13: 0013
총 네 칸 차지하고 왼쪽으로 정렬되는 13과 15: 13  15
12.1234를 소수점 둘째 자리까지만 표시하면 12.12입니다.
```

---

## Scan

콘솔을 통하여 값을 입력받을 수 있습니다.<br>

### Scan,Scanln,,Scanf

1. Scan : 공백으로 구분하여 입력
   `go run section5-scan.go`

```go
package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string
	fmt.Print("이름,성별,학교를 입력하세요")
	fmt.Scan(&name, &gen, &school)

	fmt.Println("이름", name, "성별", gen, "학교", school)
}
```

```
// result1
이름,성별,학교를 입력하세요kim male middle
이름 kim 성별 male 학교 middle

// result2
이름,성별,학교를 입력하세요kim
mail
middle
이름 kim 성별 mail 학교 middle
```

2. Scanln : 공백,개행으로 구분하여 입력
   `go run section5-scanf.go`

```go
package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string
	fmt.Print("이름,성별,학교를 입력하세요")
	fmt.Scanln(&name, &gen, &school)

	fmt.Println("이름", name, "성별", gen, "학교", school)
}
```

```
//result1
이름,성별,학교를 입력하세요kim male middle
이름 kim 성별 male 학교 middle

//rsult2
이름,성별,학교를 입력하세요kim
이름 kim 성별  학교
```

3. Scanf : 포맷팅을 이용하여 원하는 형태로 입력
   `go run section5-scanf.go`

```go
package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("-로 숫자를 나누어 2개 입력하세요")
	fmt.Scanf("%d-%d", &x, &y)

	fmt.Println("첫숫자 ", x)
	fmt.Println("뒷숫자 ", y)
}

```

```
-로 숫자를 나누어 2개 입력하세요101010101-202020202
첫숫자 101010101
뒷숫자 202020202
```

---

## 실습

### 실습1. 정돈된 표

이름 열은 전부 폭을 8로 지정하고 왼쪽 정렬을 합니다.<br>
전공학과 열은 전부 폭을 14로 지정하고 왼쪽 정렬을 합니다.<br>
학년 열은 전부 폭을 5로 지정하고 오른쪽 정렬을 합니다.<br>
모든 값은 string 형입니다.<br>
`go run section5-training1.go`

```go
package main

import "fmt"

func main() {
	fmt.Printf("%-8s%-14s%5s\n", "이름","전공학과","학년")
	fmt.Printf("%-8s%-14s%5s\n", "유현수","전자공학","3")
	fmt.Printf("%-8s%-14s%5s\n", "김윤욱","컴퓨터공학","4")
	fmt.Printf("%-8s%-14s%5s\n", "김나영","미술교육학","2")
}
```

```
이름      전공학과             학년
유현수     전자공학              3
김윤욱     컴퓨터공학             4
김나영     미술교육학             2
```

### 실습2. 신상정보 입력과 출력

주민등록번호 앞자리와 뒷자리를 저장할 int형 변수 RRNf와 RRNt을 선언합니다.<br>
이름을 저장할 string형 변수 name을 선언합니다.<br>
키를 저장할 float32형 변수 height를 선언합니다.<br>
첫 번째 줄부터 주민등록번호, 이름, 키를 입력받습니다.<br>
정보를 모두 입력하면 결괏값이 출력됩니다.<br>

`go run section5-training2.go`

```go
package main

import "fmt"

func main() {
var RRNf int
var RRNt int
var name string
var height float32
fmt.Scanf("%d-%d",&RRNf,&RRNt)
fmt.Scanf("%s",&name)
fmt.Scanf("%f",&height)
fmt.Printf("주민등록번호 앞자리는 %d, 뒷자리는 %d, 이름은 %s입니다.\n그리고 키는 %0.2f입니다.",RRNf,RRNt,name,height)
}
```

입력값

```
910101-1000000
kim
180.3345
```

결과값

```
주민등록번호 앞자리는 910101, 뒷자리는 1000000, 이름은 kim입니다.
그리고 키는 180.33입니다.
```
