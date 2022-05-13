package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := f1()
	ch2 := f2()

	data1 := <-ch1
	fmt.Println(data1)

	data2 := <-ch2
	fmt.Println(data2)

	fmt.Println("Done")
}

func f1() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "f1 completed"
	}()
	return ch
}

func f2() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- "f2 completed"
	}()
	return ch
}
