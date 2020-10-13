# 인터페이스

## 메소드 집합 인터페이스

- 인터페이스 : 메소드가 여러개 이용할때 사용을 편하게 해주기 위한 메소드를 하나로 모아놓은 것.

**구조체**는 하나의 변수에 여러개의 필요한 변수를 묶어놓은 개념입니다. 이떄 구조체의 변수들을 이용하여 연산을 하는것을 **메소드**입니다.<br>
이때 메소드도 여러개가 필요할 경우가 있습니다.<br><br>
예로 삼각형과 사각형이 있을때 넓이를 구할려면 삼각형은 밑변x높이/2 사각형은 밑변x높이로 구해야합니다. 이때 각각의 메소드를 따로 생성하고 이름을 따로 두어야 합니다. 그러다보니 프로그램의 길이가 길어지면, 해당 메소드를 관리하기 힘들어 집니다.<br><br>
인터페이스의 장점으로는 메소드들을 한눈에 보기 편리합니다.

```go
//section14-notInterface.go
package main

import (
	"fmt"
	"math" //Pi를 사용하기 위해 import함
)

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}

	fmt.Println(r1.area())
	fmt.Println(c1.area())
}
```

```
200
314.1592653589793
```

```go
//section14-interface.go
package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

	printMeasure(r1, c1, r2, c2)
}

func printMeasure(m ...geometry) { //인터페이스를 가변 인자로 하는 함수
	for _, val := range m { //가변 인자 함수의 값은 슬라이스형
		fmt.Println(val.area()) //인터페이스의 메소드 호출
	}
}
```

```
200
314.1592653589793
168
78.53981633974483
```

**notInterface**는 인터페이스를 사용하지 않고, **interface** 인터페이스를 사용했습니다.
**notInterface**는 원과 사각형의 넓이를 구하는 메소드를 각각 선언했습니다. 사각형과 원의 수식방식이 다르기 떄문입니다.<br>
출력을 하기 위해서는 개별적으로 출력이 필요합니다. 지금은 하나의 변수를 이용해서 한번만 선언하기때문에 2개의 출력문만 필요했습니다.<br>
하지만 예로 3개이상의 수식혹은 출력할 상황이 필요할때는 수 만큼 출력을 해주어야 합니다.<br>
함수를 이용하여 출력할 수 있지만, 매개변수와 형이 다르기 때문에 각각 선언해줄 필요가 있습니다<br><br>

**interface**는 인터페이스로 메소드를 묶고 출력하였습니다. 즉, `Rect`,`Circle`는 구조체의 형태가 다르지만, 인터페이스를 이용하여 개별적으로 선언이 필요가 없습니다<br>
예를들어 넓이와 둘레 의 수식을 구할떄 인터페이스의 차이입니다.

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64 // 둘레를 측정하는 메소드 추가
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perimeter() float64 { // 둘레를 측정하는 메소드 추가
	return 2 * (r.width + r.height)
}

func (c Circle) perimeter() float64 { // 둘레를 측정하는 메소드 추가
	return 2 * math.Pi * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

    printMeasure(r1, c1, r2, c2)
    // fmt.Println(r1.area())
    // fmt.Println(r1.perimeter())
    // fmt.Println(c1.area())
    // fmt.Println(c1.perimeter())
    // fmt.Println(r2.area())
    // fmt.Println(r2.perimeter())
    // fmt.Println(c2.area())
    // fmt.Println(c2.perimeter())
}

func printMeasure(m ...geometry) {
	for _, val := range m {
        fmt.Println("")
		fmt.Println(val.area())
		fmt.Println(val.perimeter())
	}
}
```

```

200
60

314.1592653589793
62.83185307179586

168
52

78.53981633974483
31.41592653589793
```

---

## 빈 인터페이스

- 인터페이스는 내용을 따로 선언하지 않아도 형으로서 사용할 수 있습니다.
- 인터페이스는 하나의 형이기 때문에 매개변수로 사용될 수 있습니다.
- 인터페이스는 **Dynamic Type** 즉,어떠한 타입도 담을 수 있는 컨테이너 입니다.
  즉, 여러개의 변수형태를 인터페이스를 이용하여 담아 사용할 수 있습니다.

```go
//section14-emptyInterface.go
package main

import "fmt"

func printVal(i interface{}) {
	fmt.Println(i)
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func main() {
	var x interface{} //빈 인터페이스 선언

	x = 1
	printVal(x)

	x = "test"
	printVal(x)

	r1 := Rect{10, 20}
	x = r1
	printVal(x)
}
```

```
1
test
{10 20}
```

### Type Assertion

인터페이스는 타입이 동적으로 변하지만, 타입을 확실하게 명시할 수 있습니다 `"변수이름.(형)"`을 이용하면 됩니다.
만약 nil일경우 에러가 발생합니다.

```go
package main

import "fmt"

func main() {
    var num interface{} = 10

    a := num
    b := num.(int)

	fmt.Printf("%T,%d\n",a,a)
    printtest(b)
}

func printtest (i interface{}){
	fmt.Printf("%T,%d\n",i,i)
}
```

```
int,10
int,10
```

---

## 실습

### 실습1 직육면체와 원기둥

```go
package main

import (
	"fmt"
	"math"
)

type cal interface {
	volume() float64
	surfaceArea() float64
}

type Cube struct {
	horizontal float64
	vertical   float64
	height     float64
}

type Cylinder struct {
	radius float64
	height float64
}

func (cu Cube) volume() float64 {
	return cu.horizontal * cu.vertical * cu.height
}

func (cu Cube) surfaceArea() float64 {
	return 2 * (cu.height*cu.horizontal + cu.height*cu.vertical + cu.vertical*cu.horizontal)
}

func (cy Cylinder) volume() float64 {
	return (cy.radius * cy.radius * math.Pi) * cy.height
}

func (cy Cylinder) surfaceArea() float64 {
	return (cy.radius*cy.radius*math.Pi)*2 + 2*cy.radius*math.Pi*cy.height
}

func printMeasure(m ...cal) {
	for _, value := range m {
		fmt.Printf("%0.2f, %0.2f\n", value.surfaceArea(), value.volume())
	}
}

func main() {
	cy1 := Cylinder{10, 10}
	cy2 := Cylinder{4.2, 15.6}
	cu1 := Cube{10.5, 20.2, 20}
	cu2 := Cube{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}
```

```
1256.64, 3141.59
522.51, 864.52
1652.20, 4242.00
724.00, 920.00
```
