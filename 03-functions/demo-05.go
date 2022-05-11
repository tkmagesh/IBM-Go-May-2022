/*
 - higher order functions
        * can be returned as return values from other functions
*/

package main

import "fmt"

func main() {
	fn := getFn()
	fn()

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	logAdd := getLogOperation(add)
	logAdd(100, 200)

	logSubtract := getLogOperation(subtract)
	logSubtract(100, 200)
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}

func getLogOperation(oper func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		result := oper(x, y)
		fmt.Println("Result =", result)
		fmt.Println("After invocation")
	}
}

/*
func logAdd(x, y int) {
	fmt.Println("Before invocation")
	add(x, y)
	fmt.Println("After invocation")
}

func logSubtract(x, y int) {
	fmt.Println("Before invocation")
	subtract(x, y)
	fmt.Println("After invocation")
}
*/

func add(x, y int) int {
	result := x + y
	return result
}

func subtract(x, y int) int {
	result := x - y
	return result
}
