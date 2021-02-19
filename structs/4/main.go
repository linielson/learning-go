package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name string `json:"Renamed"` //change the property name
	Year int `json:"-"` //does not include at json
	Color string
}

func main() {
	car := Car{"Gol", 2019, "Yellow"}
	result, _ := json.Marshal(car)
	//fmt.Println(result) ascii
	fmt.Println(string(result))
}
