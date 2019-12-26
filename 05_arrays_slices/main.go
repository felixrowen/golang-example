package main

import "fmt"

func main() {

	// Array
	fruitArr := [2]string{}

	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Mango"

	// Declare and Assign
	brandArr := [3]string{"Gucci", "Zara", "LV"}

	// Array Slice
	arrSlice := []int{23, 15, 13, 44, 9}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fmt.Println(brandArr)

	fmt.Println(arrSlice)

	// To count length of array
	fmt.Println(len(arrSlice))

	// To get range value in array, start from index 1 and stop before index 4
	fmt.Println(arrSlice[1:4])
}
