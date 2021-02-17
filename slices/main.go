package main

import "fmt"

func main() {
	// 0 - quantity items/size
	// 5 - capacity
	slice := make([]int, 0, 5)
	//slice[0] = 10// error: index out of range [0] with length 0 - was defined size 0
	//slice[1] = 11 // error: index out of range [1] with length 0 - was defined size 0
	fmt.Println(slice)

	slice = append(slice, 1, 5, 10)
	fmt.Println(slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))

	slice = append(slice, 4, 5, 6)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice)) //exceeded capacity, doubled capacity

	//forInSlice()
	pointer()
}

func forInSlice() {
	slice := make([]int, 2, 5)

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Len:", len(slice), "Cap:", cap(slice))
	}
}

func pointer() {
	slice := make([]int, 2, 5)
	sliceTest := slice // it not a copy, it is a pointer
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println(sliceTest)

	fmt.Println("----------")

	slice = append(slice, 1, 2, 3)
	slice[0] = 3
	fmt.Println(slice)
	fmt.Println(sliceTest)

	fmt.Println("---------- exceed the capacity")

	slice = append(slice, 4)
	fmt.Println(slice)
	fmt.Println(sliceTest)

	fmt.Println("---------- after exceed the capacity, they point to different arrays")
	slice[0] = 5

	fmt.Println(slice)
	fmt.Println(sliceTest)

}