package main

import "fmt"

func main() {
	//var productsRank map[string]int
	//var productsRank map[string]int = map[string]int{}
	/*
		var productsRank map[string]int = make(map[string]int)
		productsRank["Pen"] = 3
	*/
	//productsRank := map[string]int{"Pen": 2, "Pencil": 1, "Marker": 5}
	productsRank := map[string]int{
		"Pen":    2,
		"Pencil": 1,
		"Marker": 5,
	}
	fmt.Println(productsRank)

	fmt.Println("Adding a new item")
	productsRank["Stapler"] = 4
	fmt.Println(productsRank)

	fmt.Println("Iterating a map (range)")
	for product, rank := range productsRank {
		fmt.Printf("Rank of %q is %d\n", product, rank)
	}

	fmt.Println("Checking if a key exists")
	keyToCheck := "Pen"
	//keyToCheck := "Scribble-Pad"
	if rank, exists := productsRank[keyToCheck]; exists {
		fmt.Printf("Product %q exists with rank = %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Product %q does not exist\n", keyToCheck)
	}

	fmt.Println("Removing an item")
	delete(productsRank, "Pen")
	fmt.Println(productsRank)

}
