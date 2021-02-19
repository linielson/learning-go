package main

import (
	"fmt"
	car "github.com/linielson/learning-go/packages/car"
)

func main() {
	//car := car.Car{"Brasilia", "Blue"} // implicit assignment of unexported field 'color' in car.Car literal
	var car = car.Car{}
	car.Name = "Brasilia"
	fmt.Println(car.Start())
}
