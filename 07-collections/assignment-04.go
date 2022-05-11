package main

import (
	"fmt"
)

func main() {
	var primeNos = generatePrimes(3, 100)
	fmt.Println(primeNos)
}

func generatePrimes(start, end int) []int {
	primeNos := []int{}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primeNos = append(primeNos, i)
		}
	}
	return primeNos
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
