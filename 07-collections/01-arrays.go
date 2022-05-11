package main

import "fmt"

func main() {
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iterating (using indexer)
	fmt.Println("Iterating (using indexer)")
	for idx := 0; idx < len(nos); idx++ {
		no := nos[idx]
		fmt.Printf("nos[%d]=%d\n", idx, no)
	}

	//iterating (using range)
	fmt.Println("Iterating (using range)")
	for idx, no := range nos {
		fmt.Printf("nos[%d]=%d\n", idx, no)
	}

	fmt.Println("Iterating (using range) only value")
	for _, no := range nos {
		fmt.Printf("no = %d\n", no)
	}

}
