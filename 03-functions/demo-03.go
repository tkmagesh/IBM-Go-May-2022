/* higher order functions - can be assigned as values to variables
 */
package main

import "fmt"

func main() {

	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Printf("sum of 100 & 200 is %d\n", add(100, 200))

	var divide func(int, int) (int, int)
	divide = func(x, y int) (int, int) {
		return x / y, x % y
	}
	//fmt.Println(divide(100, 7))
	quotient, remainder := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d & remainder = %d\n", quotient, remainder)

}
