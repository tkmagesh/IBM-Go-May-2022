package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{}
	//emp := Employee{100, "Magesh", 10000}
	//emp := Employee{Id: 100, Name: "Magesh", Salary: 10000}
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Printf("%#v\n", emp)
	fmt.Println("EMP Name =", emp.Name)
}
