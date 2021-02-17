package main

import "fmt"

func main() {
	a := "Lini"
	switch a {
	case "Linielson":
		fmt.Println("Hey Linielson")
	case "Rosa":
		fmt.Println("Hey Rosa")
	case "Lini":
		fmt.Println("Hey Lini")
	default:
		fmt.Println("404 - Not found")
	}
}
