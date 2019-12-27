package main

import "fmt"

func main() {
	// Define Map
	emails := make(map[string]string)

	// Assign Key Value
	emails["Felix"] = "felixroo98@gmail.com"
	emails["Sharon"] = "sharon98@gmail.com"
	emails["Mike"] = "mike98@gmail.com"

	// Declare and add Key Values
	age := map[string]int{"felix": 21, "sarah": 25, "mike": 18}

	age["wilson"] = 20

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Felix"])

	delete(emails, "Sharon")
	fmt.Println(emails)
	fmt.Println(age)
}
