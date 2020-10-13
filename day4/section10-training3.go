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
	fmt.Printf("%0.2f", avg/float32(len(m)))
}
