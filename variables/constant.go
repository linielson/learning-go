package main

import "fmt"

const constWithType string = "I have a type"
const PublicConst = 10
const (
	const1 = 1
	Const2 = 2
	Const3 int = 3
	const4 float32 = 4.5
)

func funcWithConst() {
	const constInAFunc = "some value"
	fmt.Println(constInAFunc)
	fmt.Println(const1, Const2, Const3, const4, PublicConst)
}
