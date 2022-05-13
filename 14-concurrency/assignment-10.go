package main

import "fmt"

func main() {
	count := 10
	primeNosCh := genPrimes(3, count)
	for i := 0; i < count; i++ {
		fmt.Println("Prime No :", <-primeNosCh)
	}
	fmt.Println("Done")
}

func genPrimes(start, count int) <-chan int {
	ch := make(chan int)
	go func() {
		no := start
		for count > 0 {
			if isPrime(no) {
				ch <- no
				count--
			}
			no++
		}
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
