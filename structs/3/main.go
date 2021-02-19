package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name string
	Year int
	Color string
	test string // private/unexported
}

func main() {
	car := Car{"Gol", 2019, "Yellow", "private prop"}
	result, _ := json.Marshal(car)
	//fmt.Println(result) ascii
	fmt.Println(string(result))
}
