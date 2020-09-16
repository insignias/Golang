package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	swetha := person{
		firstName: "Swetha",
		lastName:  "Pappala",
		contactInfo: contactInfo{
			email:   "swetha17@gmail.com",
			zipCode: 01760,
		},
	}
	//swethaPointer := &swetha

	swetha.updateName("Shweta")
	swetha.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
