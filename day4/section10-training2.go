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
