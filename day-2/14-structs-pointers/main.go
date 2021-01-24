package main

import "fmt"

// Address of an employee
type Address struct {
	city, state string
}

func (a Address) Print() string {
	return fmt.Sprintf("%s %s", a.city, a.state)
}

// Employee refers to an employee record
type Employee struct {
	firstName string
	lastName  string
	age       int
	Address
}

// Print will display a formatted employee record
func (e Employee) Print() {
	fmt.Println("Employee Record:")
	fmt.Println("Name:", e.firstName, e.lastName)
	// Call "inherited" method with same name
	fmt.Println("Address:", e.Address.Print())
}

// updateAge will change age to the new value
func (e *Employee) updateAge(newAge int) {
	e.age = newAge
}

func main() {
	emp := Employee{
		firstName: "Peter",
		lastName:  "Parker",
		age:       22,
		Address: Address{
			city:  "New York",
			state: "New York",
		},
	}
	fmt.Println(emp)
	fmt.Println("City:", emp.city) //Printing nested field

	// Using struct method
	emp.Print()

	// Using a pointer receiver
	fmt.Println("Before:", emp.age)
	emp.updateAge(34)
	fmt.Println("After:", emp.age)
}
