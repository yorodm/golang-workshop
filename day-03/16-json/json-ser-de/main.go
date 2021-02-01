package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	City  string `json:"ciudad,omitempty"`
	State string `json:"estado,omitempty"`
}

type Employee struct {
	FirstName string `json:"nombre"`
	LastName  string `json:"apellido"`
	Age       int    `json:"edad"`
	Address
}

func main() {
	emp1 := Employee{"Peter", "Parker", 22, Address{"Manhattan", "New York"}}
	emp2 := Employee{"Hulk", "Banner", 22, Address{}}
	d, err := json.Marshal(&emp1)
	if err != nil {
		log.Fatal("No se pudo serializar")
	}
	fmt.Println(string(d))
	d, err = json.Marshal(&emp2)
	if err != nil {
		log.Fatal("No se pudo serializar")
	}
	fmt.Println(string(d))
	tony := `{"nombre":"Tony","apellido":"Stark","edad":44}`
	var emp3 Employee
	err = json.Unmarshal([]byte(tony), &emp3)
	if err != nil {
		log.Fatal("No se pudo deserializar")
	}
	fmt.Println(emp3)
}
