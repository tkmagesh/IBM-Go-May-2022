package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	ibm := &Organization{
		Name: "IBM",
		City: "Bengaluru",
	}
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    ibm,
	}
	fmt.Println(emp.Org.City)
	ibm.City = "Pune"
	fmt.Println(emp.Org.City)
}
