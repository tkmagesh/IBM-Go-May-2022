package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //creating a channel
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch //receive operation
	wg.Wait()
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result //send operation
	wg.Done()
}
