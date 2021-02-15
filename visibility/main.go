package main

import (
	"another_package"
	"fmt"
)

func main() {
	another_package.PrintPublic()
	fmt.Println(another_package.ExportedVariable)
}
