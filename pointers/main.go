package main

import "fmt"

// & print memory
// * print value

func main() {
	normalVar := 1
	fmt.Println(normalVar) // 1
	fmt.Println("Prints the address that 'normalVar' is pointing to: ", &normalVar) // 0xc0000180b8

	pointTo := &normalVar // point to address memory
	fmt.Println("Prints the 'normalVar' address too: ", pointTo) // 0xc0000180b8
	fmt.Println("Prints the 'normalVar' value: ", *pointTo) // 1
	fmt.Println("But 'pointTo' has another address: ", &pointTo) // 0xc000182020

	//pointTo := 2 // cannot use 2 (type int) as type *int in assignment

	*pointTo = 2 //change the value stored at memory address
	fmt.Println(pointTo) // 0xc0000180b8
	fmt.Println(normalVar) // 2
	fmt.Println(*pointTo) // 2

	var withType *int = &normalVar //with type

	fmt.Println(withType) // 0xc0000180b8
	fmt.Println(*withType) // 2

	fmt.Println("Tests with function")
	someVar := 10
	fmt.Println(someVar) // 10

	fmt.Println("By value")
	fmt.Println(thisIsAFunctionByValue(someVar)) // 11
	fmt.Println(someVar) // 10

	fmt.Println("By Ref")
	//thisIsAFunctionByRef(someVar) // does not work
	fmt.Println(thisIsAFunctionByRef(&someVar)) // 12
	fmt.Println(someVar) // 12

	fmt.Println("Changes directly ")
	someVar = thisIsAFunctionByValue(someVar)
	fmt.Println(someVar) // 13

	fmt.Println("Changes by procedure")
	//thisIsLikeAProcedure(someVar) // does not work - cannot use someVar (type int) as type *int in argument to thisIsLikeAProcedure
	//fmt.Println(thisIsLikeAProcedure(&someVar)) // thisIsLikeAProcedure(&someVar) used as value
	thisIsLikeAProcedure(&someVar)
	fmt.Println(someVar) // 16
}

//receive an Int value and return an Int value too
func thisIsAFunctionByValue (value int) int {
	value = value + 1
	return value //function needs "return"
}

// * only accept the address of a int variable
func thisIsAFunctionByRef (value *int) int {
	*value = *value + 2
	return *value
}

func thisIsLikeAProcedure (value *int) {
	*value = *value + 3
}
