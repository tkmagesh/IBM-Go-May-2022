package models

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (e Employee) WhoAmI() {
	fmt.Println("I am an Employee - ", e.Name)
}
