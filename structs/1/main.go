package main

import fmt "fmt"

//New type, can be used into any function
//We can't create a primitive type
//type SuperNumber int

//We can create any type, but they are used individual
//type CarName string
//type CarYear int
//type Color string

type Car struct {
	Name string
	Year int
	Color string
}

func main() {
    //x := []SuperNumber //== slice []int
	car1 := Car{"Corolla", 2017, "White"}
	car2 := Car{"BMW", 2017, "Black"}
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.Name, car1.Year, car1.Color)
	fmt.Println(Car.info(car2))
	fmt.Println(car1.info())
}
     //attach func info at variable c with type Car
     //return string
func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}
