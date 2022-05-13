package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan bool)
	primeNosCh := genPrimes(stopCh)

	go func() {
		var input string
		fmt.Scanln(&input)
		stopCh <- true
	}()

	fmt.Println("Hit ENTER to stop....")
	for primeNo := range primeNosCh {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(stopCh <-chan bool) <-chan int {
	ch := make(chan int)

	go func() {
		no := 3
	LOOP:
		for {
			if !isPrime(no) {
				no++
				select {
				case <-stopCh:
					break LOOP
				default:
					continue LOOP
				}
			}
			select {
			case ch <- no:
				time.Sleep(500 * time.Millisecond)
				no++
			case <-stopCh:
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
