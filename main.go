package main

import (
	"fmt"

	"datatypes.com/organization"
)

func main() {
	p := organization.NewPerson("John", "Smith", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	err := p.SetTwitterHandler("@helloworld")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("Error setting Twitter handler %s\n", err.Error())
	}

	name1 := Name{First: "John", Last: "Smith"}
	name2 := Name{First: "John", Last: "Smith"}

	if name1 == name2 {
		println("Names match")
	}

	// name3 := OtherName{First: "John", Last: "Smith"}

	// not possible
	// if name1 == name3 {
	// }
}

// variables whose types have same memory layout can be compared
// However, slices and methods can't be compared
type Name struct {
	First string
	Last  string
}

// variables whose types have same fields, but are of different types can't be compared
type OtherName struct {
	First string
	Last  string
}

// Two interfaces that have the same fields, methods and parameters can be compared
