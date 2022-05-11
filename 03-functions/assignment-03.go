package main

import "fmt"

func main() {
	for {
		userChoice := getUserChoice()
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		result := doOperation(userChoice)
		fmt.Println("Result =", result)
	}
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Enter your choice")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter 2 numbers")
	fmt.Scanln(&n1, &n2)
	return n1, n2
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func doOperation(userChoice int) int {
	n1, n2 := getOperands()
	result := 0
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	return result
}
