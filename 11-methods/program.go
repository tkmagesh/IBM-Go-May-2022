package main

import (
	"fmt"
	"methods-demo/models"
)

type MyStr string

func (myStr MyStr) Length() int {
	return len(myStr)
}

func main() {
	e := models.Employee{Name: "Magesh"}
	e.WhoAmI()

	//type conversion
	var i int = 100
	var f float32
	f = float32(i)
	fmt.Println(f)

	s := MyStr("This is a sample string")
	fmt.Println(s.Length())
}
