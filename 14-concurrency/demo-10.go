package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("[@goroutine] attempting to send data")
		ch <- 100
		fmt.Println("[@goroutine] successfully sent data")
	}()
	fmt.Println("[@main] attempting to receive data")
	data := <-ch
	fmt.Println("[@main] successfully received data")
	fmt.Println("data from channel = ", data)
}
