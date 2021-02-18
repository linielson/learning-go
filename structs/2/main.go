package main

import fmt "fmt"

type Car struct {
	Name string
	Year int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name string //overwrite the name
}

func main() {
	sCar := SuperCar{
		Car: Car{
			"Fusca", 2030, "Blue",
		},
		CanFly: true,
		Name: "Celta Overwrite",
	}

	fmt.Println("-------- Super Car 00 --------")
	fmt.Println(sCar)
	fmt.Println("Name overwrite:", sCar.Name)
	fmt.Println("Name of Car:", sCar.Car.Name)
	fmt.Println(sCar.info())

	fmt.Println("-------- Super Car 01 --------")
	car1 := Car{"Corolla", 2017, "White"}
	sCar1 := SuperCar{
		Car: car1,
		CanFly: false,
	}
	fmt.Println(sCar1.info())

}
func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}
