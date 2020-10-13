package main

import "fmt"

var globalA = 5 //함수 밖에서는 'var' 키워드를 입력해야함.
				// 꼭 형을 명시하지 않아도 됨
func main() {
    var a string = "String value"
    fmt.Println(a)

    var b int = 10
    fmt.Println(b)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "short assignment value"
    fmt.Println(f)
	
	fmt.Println(globalA)

	var g, h int = 10, 20
    fmt.Println(g, h)

	i, j, k := 1, 2, 3
    fmt.Println(i, j, k)

    var str1, str2 string = "Decalere", "more than 2"
    fmt.Println(str1, str2)
}