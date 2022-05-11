package main

import "fmt"

func main() {
	var no int = 10

	var noPtr *int = &no //address from value
	var x int = *noPtr   // value from address (dereferencing)
	//no == *(&no)
	fmt.Println(no, noPtr, x)

	fmt.Printf("Before incrementing, no = %d\n", no) //=> 10
	increment(&no)
	fmt.Printf("After incrementing, no = %d\n", no) //=> 11

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) {
	*x += 1
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
