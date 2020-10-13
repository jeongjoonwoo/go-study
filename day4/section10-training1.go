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
