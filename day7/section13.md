# 구조체와 메소드

## 구조체

구조체 : 하나의 변수에 여러개의 자료형을 정의하는 custom data type.

- 특정 구조의 변수가 여러번 선언될때, 해당 변수를 매번 선언하면 가독성이 떨어지므로, 하나의 구조를 선언한 후 지속적으로 이용하는 형태입니다.
- Go는 객체지향언어이므로 클래스,객체,상속의 개념이 존재하지 않습니다.

```go
//구조선언
type 구조체명 struce {
    변수명1 변수타입1
    변수명2 변수타입2
    변수명3 변수타입3
    변수명4 변수타입4
}
```

```go
// 생성
객체이름 := 구조체이름{저장할값} //구조체
객체이름 = new(구조체이름) //포인터구조체
```

- 구조체는 구조만 선언한 형식이므로, 실제로 사용하기 위해선 별도의 선언이 필요합니다.
- `객체이름.필드이름=저장할값`으로 저장할 수 있습니다.
- 필드가 생력된 경우 필드값은 Zero Value를 갖습니다.

```go
//section13-struct.go
package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000000000"
	fmt.Println(p1)

	p2 := person{"nam", 31, "01022220000"} // 필드 이름을 생력할 시 순서대로 저장함
	fmt.Println(p2)

	p3 := person{contact: "01011110000", name: "park", age: 23} // 필드 이름을 명시할 시 순서와 상관 없이 저장할 수 있음
	fmt.Println(p3)

	p3.name = "ryu" //필드에 저장된 값을 수정할 수 있음
	fmt.Println(p3)

	fmt.Println(p3.contact) //필드 값의 개별 접근도 가능함
}
```

```
{ 0 }
{kim 25 01000000000}
{nam 31 01022220000}
{park 23 01011110000}
{ryu 23 01011110000}
01011110000
```

- 선언후 초기화를 하지 않았을때(p1)
  - 출력시 0값이 출력
  - 변수명.필드명으로 값을 입력가능
- 선언 후 초기화시 필드명을 입력하지 않았을때(p2)
  - 순서대로 입력
- 선언 후 필드명과 함께 선언하였을때(p3)
  - 필드명에 따라 값이 입력
- 필드명으로 값을 변경하였을때
  - 값이 변경 되어짐.

### 구조체 역참조 방법

1. new(구조체이름)
2. 구조체이름앞에 & 붙이기

다른 변수는 역참조(변수의 주소에 접근)을 위하여 \*를 이용했지만, 구조체는 자동으로 역참조 됩니다.

```go
//section13-name.go
package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func addAgeRef(a *person) { //Pass by reference
	a.age += 4 //자동 역참조 * 생략
}

func addAgeVal(a person) { //Pass by value
	a.age += 4
}

func main() {
	var p1 = new(person) //포인터 구조체 객체 생성
	var p2 = person{}    // 빈 구조체 객체 생성

	fmt.Println(p1, p2)

	p1.age = 25
	p2.age = 25

	addAgeRef(p1) //&을 쓰지 않아도 됨
	addAgeVal(p2)

	fmt.Println(p1.age)
	fmt.Println(p2.age)

	addAgeRef(&p2) //&을 붙이면 pass by reference 가능
	fmt.Println(p2.age)
}
```

```
&{ 0 } { 0 }
29
25
29
```

### 생성자

구조체는 사용하기 이전에 타입으로 구조와 변수형을 선언해주어야 합니다.
맵 형태일 경우 구조체로 선언할때, 해당 맵 필드도 같이 초기화 해야하는 번거로움을 해결할 수 있습니다.

```go
type mapStruct struct{ //맵 형태의 data필드를 가지는 "mapStruct"를 정의함
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{} //구조체 객체를 생성하고 초기화함
	d.data = map[int]string{} //data 필드의 맵을 초기화함
	return &d //초기화 한 포인터 구조체를 반환함
}
```

생성자 함수는 구조체 객체를포인터와 함께 반환합니다.
객체를 생성하는 생성자를 만들려면 반환할 경우 포인터 연산자 없이 사용할 수 있습니다.

```go
package main

import "fmt"

type mapStruct struct {
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}
	d.data = map[int]string{}
	return &d
}
func newStruct2() mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}
	d.data = map[int]string{}
	return d
}

func main() {
	s1 := newStruct() // 생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}
	s2.data = map[int]string{}
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)

	s3 := newStruct() // 생성자 호출
	s3.data[1] = "hello"
	s3.data[10] = "world"

	fmt.Println(s3.data[1])
}
```

```
&{map[1:hello 10:world]}
{map[1:hello 10:world]}
hello
```

---

## 메소드

- 메소드 : 특성 속성들의 기능을 수행하기 위해 만들어진 특변할 함수
  언어를 이용하다보면 sort,push,pop,out 등등 언어별로 선언하는 메소드들이 있습니다. 이를 정의하고 사용하는 방식입니다.

```go
package main

import "fmt"

type triangle struct {
	width, height float32
}

func triArea(s *triangle) float32 { //'new'로 생성한 구조체 객체는 포인터값 반환
	return s.width * s.height / 2 //포인터 구조체는 자동 역참조 "*" 생략
}

func (s triangle) triArea() float32 { //value receiver
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea := triArea(tri1)
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea)

	triarea2 := tri1.triArea()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea2)
}
```

간단하게 삼각형의 넓이를 구하는 함수입니다.

1. `func triArea(s *triangle) float32 `
   - 함수를 선언하여 사용합니다.
   - 하지만 이 방식에서 단순하게 매개변수로 던져주어 구조체를 이용해야할 필요를 느끼지 못합니다.
2. `func (s triangle) triArea() float32 `
   - 메소드 형식입니다.
   - 구조체를 이용하여 선언이나 전달방식이 간단합니다.

### Value Receiver와 Pointer Receiver

- value Receiver : 구조체의 값을 복사후 반환
- pointer Receiver : 구조체의 필드값의 주소를 전달

```go
package main

import "fmt"

type triangle struct {
	width, height float32
}

func (s triangle) triAreaVal() float32 { //Value Receiver
	s.width += 10
	return s.width * s.height / 2
}

func (s *triangle) triAreaRef() float32 { //Pointer Reciver
	s.width += 10
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea_val := tri1.triAreaVal()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea_val)

	triarea_ref := tri1.triAreaRef()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea_ref)
}
```

```
삼각형 너비:12.50, 높이:5.20 일때, 넓이:58.50
삼각형 너비:22.50, 높이:5.20 일때, 넓이:58.50
```

---

## 실습

### 실습1 성적저장프로그램

```go
package main

import "fmt"

type student struct {
	name      string
	sex       string
	scoreData map[string]int
}

func newScore() *student {
	studentData := student{}
	studentData.scoreData = map[string]int{}
	return &studentData
}



func main() {
	var stuNum, subNum, score int
	var name, sex, subject string

	fmt.Scanln(&stuNum, &subNum)

	s := make([]student, stuNum)


	for i := 0; i < stuNum; i++ {
		fmt.Scanln(&name, &sex)
		student := newScore()
		student.name = name
		student.sex = sex
		for j := 0; j < subNum; j++ {
			fmt.Scanln(&subject, &score)

			student.scoreData[subject] = score
		}
		s[i] = *student
	}

	for i:=0;i<stuNum;i++ {
 		fmt.Println("----------")
		fmt.Println(s[i].name,s[i].sex)

 		for index, val := range s[i].scoreData {
 			fmt.Println(index, val)
 		}

 	}
	fmt.Println("----------")
}
```

### 실습2

```go
package main

import "fmt"

const g = 9.8

type mul struct {
	m  float32
	v  float32
	h  float32
	ke float32
	pe float32
	me float32
}

func (data *mul) keMul() float32 {
	return data.m * data.v * data.v / 2
}

func (data *mul) peMul() float32 {
	return data.m *g *data.h
}

func main() {
	var objectNum int
	var m, v, h float32

	fmt.Scanln(&objectNum)

	object := make([]mul, objectNum)

	for i := 0; i < objectNum; i++ {
		fmt.Scanln(&m, &v, &h)
		object[i].m = m
		object[i].v = v
		object[i].h = h
		object[i].ke = object[i].keMul()
		object[i].pe = object[i].peMul()
		object[i].me = object[i].ke+ object[i].pe
	}

	for _, o := range object {
		fmt.Println(o.ke, o.pe, o.me)
	}

}
```
