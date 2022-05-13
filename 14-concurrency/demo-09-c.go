package main

import (
	"fmt"
)

func main() {
	fmt.Println("main started")
	ch := add(100, 200)
	result := <-ch //receive operation
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result //send operation
	}()
	return ch
}
