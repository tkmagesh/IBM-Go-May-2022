package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

//Communicate by sharing memory  //=> NOT advisable
var counter int64

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

	atomic.AddInt64(&counter, 1)
	wg.Done() //decrementing the counter by 1
}
