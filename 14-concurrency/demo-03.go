package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(5000)
	countStr := os.Args[1]
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("Usage: demo-03 [count]")
		return
	}
	fmt.Println("main started")

	for i := 1; i <= count; i++ {
		wg.Add(1) // => initializing the counter
		go fn(i)
	}

	wg.Wait() //blocked until the counter becomes 0
	fmt.Println("main completed")

	var input string
	fmt.Scanln(&input)
}

func fn(id int) {
	fmt.Println("fn started -", id)
	duration := time.Duration(uint64(rand.Intn(5000)) * uint64(time.Millisecond))
	time.Sleep(duration)
	wg.Done() //decrementing the counter by 1
}
