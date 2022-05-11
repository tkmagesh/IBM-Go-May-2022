/*
 - higher order functions
        * can be passed as arguments to other functions
*/

package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	exec(fn)

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func fn() {
	fmt.Println("fn invoked")
}

func exec(f func()) {
	f()
}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("Before invocation")
	oper(x, y)
	fmt.Println("After invocation")
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

func add(x, y int) {
	result := x + y
	fmt.Println("Result =", result)
}

func subtract(x, y int) {
	result := x - y
	fmt.Println("Result =", result)
}
