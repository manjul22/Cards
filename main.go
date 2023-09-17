package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v",alex)

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@email.com",
			zipCode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v",p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
