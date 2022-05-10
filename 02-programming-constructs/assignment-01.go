package main

import "fmt"

func main() {
	for no, count := 3, 0; count < 100; {
		isPrime := true
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d is prime\n", no)
			count++
		}
		no++
	}
}
