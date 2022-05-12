package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling for execution
	f2()

	//time.Sleep(5 * time.Millisecond)
	var input string
	fmt.Scanln(&input)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
