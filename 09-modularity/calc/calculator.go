package calc

import "fmt"

var opCount int

func init() {
	fmt.Println("calc package initialized")
}

func OpCount() int {
	return opCount
}
