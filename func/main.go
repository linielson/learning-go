package main

import "fmt"

func main() {
	x := funcName(5)
	fmt.Println(x)
	fmt.Println(PublicFunction("This is"), "Man")
	fmt.Println(namedReturn("Named Return is not common"))
	fmt.Println(moreReturn("Linielson", "Rosa"))
	name, lastName := moreReturn("Linielson", "Rosa")
	fmt.Println(name)
	fmt.Println(lastName)
	fmt.Println(variadicFunc(1, 2, 3, 4, 5))
	anonymousFuncExample()

	resInsideFunc := funcInsideFunc()()
	fmt.Println(resInsideFunc)
	insideFunc := funcInsideFunc()
	fmt.Println(insideFunc())

	//FIXME closure
}

func funcName(a int) int {
	return a * a
}

func PublicFunction(a string) string {
	return a + " X"
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturn(a, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int
	for _, value := range x { //index, value
		res += value
	}
	return res
}

func anonymousFuncExample() {
	value := 0
	anonymousAdd := func() int {
		value += 2
		return value
	}

	fmt.Println(anonymousAdd())
	fmt.Println(anonymousAdd())
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}