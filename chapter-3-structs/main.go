package main

import "fmt"

type typeContact struct{
	email string
	zip int
}

type typePerson struct {
	firstName string
	lastName string
	contact typeContact
}

func main () {
	var alex typePerson
	alex.firstName = "Alex"
	alex.lastName = "Smith"
	alex.contact = typeContact {
		email: "sidd@gmail.com",
	}
	alex.contact.zip = 11223
	alex.print()

	jim := typePerson{
		firstName: "Jim", 
		lastName: "Beam",
		contact: typeContact {
			email: "jum@hh.com",
			zip: 12234,
		},
	}
	jim.print()
	
}

func (p typePerson) print() {
	fmt.Printf("%+v", p)
}