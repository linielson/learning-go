package main

import (
	"fmt"
	"github.com/linielson/learning-go/another_package"
)

func main() {
	another_package.PrintPublic()
	fmt.Println(another_package.ExportedVariable)
}
