package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var counter int // data race

func main() {
	rand.Seed(5000)
	count := 100
	fmt.Println("main started")

	for i := 1; i <= count; i++ {
		wg.Add(1) // => initializing the counter
		go fn(i)
	}

	wg.Wait() //blocked until the counter becomes 0
	fmt.Println("main completed")

	fmt.Println("counter = ", counter)
}

func fn(id int) {
	fmt.Println("fn started -", id)
	duration := time.Duration(uint64(rand.Intn(5000)) * uint64(time.Millisecond))
	time.Sleep(duration)
	counter++
	wg.Done() //decrementing the counter by 1
}
