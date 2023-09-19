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

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@email.com",
	// 		zipCode: 94000,
	// 	},
	// }

	// jim.updateName("jimmy")
	// jim.print()

	colors := map[string]string {
		"red":"#ff0000",
		"green": "#4f745",
		"white": "#ffffff",
	}
	//var colors map[string]string

	//colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// 	delete(colors,"white")
	//fmt.Println(colors)
	printMap(colors)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func printMap(c map[string]string) {
	for color,hex := range c {
		fmt.Println("Hex code for", color, "is",hex)
	}
}
