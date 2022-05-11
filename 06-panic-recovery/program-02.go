package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := appDivide(100, 0)
	if err != nil {
		fmt.Println("Unable to divide. Try something else")
	} else {
		fmt.Println("Result =", result)
	}
}

func appDivide(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("unable to divide")
			return
		}
	}()
	result = divide(x, y)
	return
}

func divide(x, y int) (result int) {
	result = x / y
	return
}
