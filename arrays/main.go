package main

import "fmt"

func main() {
	var ar [10]int
	fmt.Println(ar)
	fmt.Println(len(ar))
	ar[0] = 8
	ar[1] = 6
	ar[len(ar)-1] = 12
	fmt.Println(ar)
	fmt.Println(ar[5])

	var arStr [5]string
	fmt.Println(arStr)

	initializedAr := [5]int { 3, 6, 9, 2 }
	fmt.Println(initializedAr)
}
