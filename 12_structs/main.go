package main

import (
	"fmt"
	"strconv"
)

// Person struct is to define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Felix", lastName: "Rowen", city: "Jakarta", gender: "m", age: 21}
	// Alternative
	person2 := Person{"Sarah", "Witty", "Boston", "f", 22}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Williams")
	person2.hasBirthday()
	fmt.Println(person2.greet())
}
