package main

import "fmt"

func main() {

	if 4/2 == 1 { //괄호없이
		fmt.Print("first")
	} else if 5/2 == 0 {
		fmt.Println("second")
	} else {
		fmt.Println("last")
	}

}
