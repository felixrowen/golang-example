package main

import "fmt"

func main() {

	name, email := "Felix Rowen", "felixroo98@gmail.com"

	age := 21

	isCool := true
	isCool = false

	size := 7.8
	// var size float32 = 5.6

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T \n", size)
}
