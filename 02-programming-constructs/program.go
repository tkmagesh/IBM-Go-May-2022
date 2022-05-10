package main

import "fmt"

func main() {
	//if statement
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 { //=> 'no' variable scope is limited to the 'if' block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	//for (v1.0)
	fmt.Println()
	fmt.Println("for-loop (regular)")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	/* for (v2.0) */
	fmt.Println("for-loop (while)")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	/* for (v3.0) */
	fmt.Println("for-loop (infinite)")
	no := 1
	for {
		if no > 100 {
			break
		}
		no += no
	}
	fmt.Printf("no = %d\n", no)
}
