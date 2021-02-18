package main

import "fmt"

func main() {
	m := make(map[string]int) //index [String] / value: Int
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)
	delete(m, "c")
	fmt.Println(m)
	fmt.Println(m["c"]) //return 0, because is int
	m["b"] = 15
	fmt.Println(m)

	value, exists := m["c"]
	fmt.Println("Does m['c'] exist?", exists)
	fmt.Println("m['c'] value:", value)

	value, exists = m["a"]
	fmt.Println("Does m['a'] exist?", exists)
	fmt.Println("m['a'] value:", value)

	var m2 = map[string]int{}
	fmt.Println("Empty map created without make", m2)

	newMap := map[string]int{ "x":5, "y":10 }
	fmt.Println("new map populated", newMap)

	if value, exists := m["c"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("Ops..")
	}

}
