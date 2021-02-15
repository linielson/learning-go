package main

import "fmt"

	var y string = "Hello"
	var w, z string = "H", "W"

func main()  {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `
		a
		b
		c
	`

	var x int
	x = 22

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", y)
	fmt.Printf("%v \n", w)
	fmt.Printf("%v \n", z)
	fmt.Printf("%v \n", x)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", w)
	fmt.Printf("%T \n", z)
	fmt.Printf("%T \n", x)

}
