package main

import "fmt"

func main() {

	var nos []int = make([]int, 5, 10)
	nos[0] = 3
	nos[1] = 1
	nos[2] = 4
	nos[3] = 2
	nos[4] = 5
	fmt.Println(nos)

	//append
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	//var products = []string{}
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	var products = make([]string, 0, 10)
	products = append(products, "Pen", "Pencil", "Marker")
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	products = append(products, "Mouse")
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	products = append(products, "Scribble-Pad", "Sharpner")
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)

	products = append(products, "Stappler")
	fmt.Printf("len(products) = %d, cap(products) = %d, products = %v\n", len(products), cap(products), products)
}
