package main

import "fmt"

// Permanent represent an employee with a permanent position
type Permanent struct {
	empID    int
	basicPay int
	pension  int
}

// Contract represent an employee with a contractual position
type Contract struct {
	empID    int
	basicPay int
}

func (p Permanent) calculateSalary() int {
	return p.basicPay + p.pension
}

func (p Contract) calculateSalary() int {
	return p.basicPay
}

// SalaryCalculator calculates salary
type SalaryCalculator interface {
	calculateSalary() int
}

func main() {
	sc := [...]SalaryCalculator{Permanent{1, 5000, 50}, Contract{3, 7000}}
	totalPayout := 0
	for _, v := range sc {
		totalPayout += v.calculateSalary()
	}
	fmt.Println("Total Salary:", totalPayout)
}
