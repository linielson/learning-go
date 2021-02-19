package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name string
	Year int
	Color string
}

func main() {
	var car Car

	//slice of bytes
	data := []byte(`{ "Name":"Gol", "Year": 2016, "Color": "Black" }`)

	json.Unmarshal(data, car)
	fmt.Println(car.Name, car.Year, car.Color) // 0

	//need to use pointer
	json.Unmarshal(data, &car)
	fmt.Println(car.Name, car.Year, car.Color) // Gol 2016 Black

	ignoreUnexpectProperty()
	missingProperty()
}

func ignoreUnexpectProperty() {
	var car Car
	data := []byte(`{ "Name":"Gol", "Year": 2016, "Other": "Property" }`)

	json.Unmarshal(data, &car)
	fmt.Println(car.Name, car.Year, car.Color)
	//fmt.Println(car.Name, car.Year, car.Color, car.Other) // car.Other undefined (type Car has no field or method Other)
}

func missingProperty() {
	var car Car
	data := []byte(`{ "Name":"Gol", "Year": 2016 }`)

	json.Unmarshal(data, &car)
	fmt.Println(car.Name, car.Year, car.Color)
}

