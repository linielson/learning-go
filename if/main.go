package main

import "fmt"

func main() {
	a := 10
	if a > 10 {
		fmt.Println("A > than 10")
	} else if a < 15 {
		fmt.Println("A < than 15")
	} else {
		fmt.Println("Not expected")
	}

	b := true

	if b {
		fmt.Println("B is true")
	}

	if x := "attr conditional"; b {
		fmt.Println(x) // "attr conditional"
	}
}
