package main

import (
	"fmt"
)

type person struct {
	firstName string
}

func (p *person) name() string {
	return p.firstName
}

type employee struct {
	EmployeeID string
	person
}

func main() {
	empl := employee{}
	empl.firstName = "Milena"
	fmt.Printf("directly: %s\n", empl.firstName)
	fmt.Printf("inderectly: %s\n", empl.person.firstName)

	fmt.Println("does not matter")
	empl.person.firstName = "Alperen"
	fmt.Printf("directly: %s\n", empl.firstName)
	fmt.Printf("inderectly: %s\n", empl.person.firstName)

}
