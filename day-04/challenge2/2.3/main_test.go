package main

import "testing"

func TestNetIncome(t *testing.T) {
	passengers := [...]Passenger{&BasePassenger{}, &LastMinutePassenger{}, &EmployeeAirlinePassenger{}, &EmployeeAirlineLastMinutePassenger{}}
	var in float32 = 1000
	var want float32 = 2000

	got := NetIncome(passengers[:], in)

	if got != want {
		t.Errorf("netIncome(%f) == %f, want %f", in, got, want)
	}
}
