package main

import "fmt"

type contact struct {
	email    string
	postcode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contact{
			email:    "jim@gmail.com",
			postcode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
