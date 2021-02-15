package main

import "fmt"

var packageScope string = "global scope 1"

func main() {
	funcScope := 10
	fmt.Println(funcScope)
	fmt.Println(packageScope)
	anotherFunction()
	fmt.Println(anotherFile)
}

func anotherFunction() {
	//fmt.Println(funcScope) error, does not work, because is from another scope
	fmt.Println(packageScope)
	fmt.Println(packageScope2)
	funcFromAnotherFile()
}

var packageScope2 string = "global scope 2"
