package main

import "testing"

func TestNetIncome(t *testing.T) {
	passengers := [...]Passenger{&BasePassenger{}, &LastMinutePassenger{}, &EmployeeAirlinePassenger{}}
	var in float32 = 1000
	var want float32 = 1500

	got := NetIncome(passengers[:], in)

	if got != want {
		t.Errorf("netIncome(%f) == %f, want %f", in, got, want)
	}
}
