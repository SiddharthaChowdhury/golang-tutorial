package main

import "fmt"

type typePerson struct {
	firstName string
	lastName string
}

func main () {
	// Pointer
	var jim typePerson
	jimPointer := &jim
	jim.firstName = "Jim"
	jimPointer.lastName = "Smith"
	jimPointer.update("Jimmy")
	jim.print()

	// Pointer shortcut
	var alex typePerson
	alex.firstName = "Alex"
	alex.lastName = "Doe"
	alex.update("Alexa")
	alex.print()
}

func (p *typePerson) update (firstName string) {
	(*p).firstName = firstName
}

func (p typePerson) print () {
	fmt.Printf("%+v", p)
}

/*
	Following dataTypes needs POINTER to update/change value when passed to a function (ie. update value in another function)

	1.	int
	2.	float
	3.	string
	5.	bool
	6.	structs
*/