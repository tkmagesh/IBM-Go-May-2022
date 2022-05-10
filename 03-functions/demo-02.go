/* anonymous functions */
package main

import "fmt"

func main() {

	func() {
		fmt.Println("fn invoked")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	greetMsg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}("Suresh")
	fmt.Println(greetMsg)

	quotient, remainder := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
}
