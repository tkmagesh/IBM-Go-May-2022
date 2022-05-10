package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World!"
		fmt.Println(msg)
	*/

	/*
		var msg string = "Hello World!"
		fmt.Println(msg)
	*/

	//type inference
	//var msg = "Hello World!"

	msg := "Hello World!" // => := syntax is applicable ONLY in a function (NOT at the package level)
	fmt.Println(msg)

	//multiple variable declarations & initializations
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		str = "Sum of x and y is"
		result = x + y
		fmt.Println(str, result)
	*/
	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of x and y is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of x and y is"
		result = x + y
	*/

	/*
		var (
			x, y, result int    = 100, 200, 0
			str          string = "Sum of x and y is"
		)
	*/
	/*
		var (
			x, y, result = 100, 200, 0
			str          = "Sum of x and y is"
		)
	*/
	/*
		var x, y, result, str = 100, 200, 0, "Sum of x and y is"
		result = x + y
	*/

	/*
		x, y, result, str := 100, 200, 0, "Sum of x and y is"
		result = x + y
	*/
	x, y, str := 100, 200, "Sum of x and y is"
	result := x + y
	fmt.Println(str, result)
}
