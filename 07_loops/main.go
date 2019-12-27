package main

import "fmt"

func main() {
	// Long Method
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short Method
	for j := 0; j <= 10; j++ {
		fmt.Printf("Number %d\n", j)
	}

	// FizzBuzz
	for k := 1; k <= 100; k++ {
		if k%3 == 0 {
			fmt.Println("Fizz")
		} else if k%5 == 0 {
			fmt.Println("Buzz")
		} else if k%15 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(k)
		}
	}
}
