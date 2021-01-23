package main

import "fmt"

type Passenger interface {
	GetDiscount() float32
}

type BasePassenger struct {
}

func (*BasePassenger) GetDiscount() float32 {
	return 0
}

type LastMinutePassenger struct {
}

func (*LastMinutePassenger) GetDiscount() float32 {
	return 0.5
}

type EmployeeAirlinePassenger struct {
}

func (*EmployeeAirlinePassenger) GetDiscount() float32 {
	return 1
}

type EmployeeAirlineLastMinutePassenger struct {
	LastMinutePassenger
	EmployeeAirlinePassenger
}

func (elm *EmployeeAirlineLastMinutePassenger) GetDiscount() float32 {
	return elm.LastMinutePassenger.GetDiscount() + elm.EmployeeAirlinePassenger.GetDiscount()
}

func NetIncome(passengers []Passenger, baseTicketPrice float32) float32 {
	var total float32
	for _, p := range passengers {
		total += baseTicketPrice - p.GetDiscount()*baseTicketPrice
	}
	return total
}

func main() {
	passengers := [...]Passenger{&BasePassenger{}, &LastMinutePassenger{}, &EmployeeAirlinePassenger{}, &EmployeeAirlineLastMinutePassenger{}}
	fmt.Println(NetIncome(passengers[:], 1000))
}
