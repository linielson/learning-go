package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

type motorcycle struct {
	name string
}

func startVehicle(v vehicle) string {
	return v.start()
}

func (c car) start() string {
	return "The car "	+ c.name + " has been started"
}

func (m motorcycle) start() string {
	return "The motorcycle "	+ m.name + " has been started"
}

func main() {
	c := car{"Opala"}
	m := motorcycle{"Mobilete"}

	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))
	//fmt.Println(c.startCar())
	//fmt.Println(m.startMotorcycle())
}


