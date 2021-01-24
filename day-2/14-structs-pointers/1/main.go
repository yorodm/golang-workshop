package main

import "fmt"

// Address of an employee
type Address struct {
	city, state string
}

// Employee refers to an employee record
type Employee struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

// Print will display a formatted employee record
func (e Employee) Print() {
	fmt.Println("Employee Record:")
	fmt.Println("Name:", e.firstName, e.lastName)
	fmt.Println("Address:", e.address)
}

// updateAge will change age to the new value
func (e *Employee) updateAge(newAge int) {
	e.age = newAge
}

func main() {
	emp := Employee{
		firstName: "Allen",
		lastName:  "Varghese",
		age:       33,
	}

	// Adding details for the nested struct
	emp.address = Address{
		city:  "AA",
		state: "CO",
	}

	fmt.Println(emp)
	fmt.Println("City:", emp.address.city) //Printing nested field

	// Using struct method
	emp.Print()

	// Using a pointer receiver
	fmt.Println("Before:", emp.age)
	emp.updateAge(34)
	fmt.Println("After:", emp.age)
}
