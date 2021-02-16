package main

import "fmt"

func main() {
	simpleFor()
	simpleForDesc()
	simpleWhile()
	//infiniteLoop()
}

func simpleFor()  {
	fmt.Println("FOR")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func simpleForDesc()  {
	fmt.Println("DESC FOR")
	for i := 5; i >= 1; i-- {
		fmt.Println(i)
	}
}

func simpleWhile() {
	fmt.Println("WHILE")
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}

func infiniteLoop() {
	x := 0
	//for true {
	for {
		x++

		if x == 50 {
			continue
		}

		if x == 100 {
			break
		}

		fmt.Println(x)
	}
}
