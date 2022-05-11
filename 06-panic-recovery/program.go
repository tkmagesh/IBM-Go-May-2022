package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong!")
			fmt.Println(err)
			return
		}
		fmt.Println("Thank you!!")
	}()
	result := divide(100, 0)
	fmt.Println("Result =", result)
}

func divide(x, y int) (result int) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		panic(err)
	}
	result = x / y
	return
}
