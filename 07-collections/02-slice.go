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

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo :] => from lo to len-1
		[:hi] => from 0 to hi - 1
		[:] => copy of the slice
	*/

	fmt.Println("Slicing")
	fmt.Println("products[2:4] => ", products[2:4])
	fmt.Println("products[:4] => ", products[:4])
	fmt.Println("products[2:] => ", products[2:])

	newProducts := products[2:4]
	newProducts[0] = "Sketch-Pen"
	fmt.Println("newProducts => ", newProducts)
	fmt.Println("products => ", products)

	x := []string{}
	x = append(x, products[:3]...)
	x = append(x, products[4:]...)
	fmt.Println(x)

	//creating a copy of the slice
	nos1 := []int{10, 20, 30, 40}
	nos2 := make([]int, len(nos1))
	copy(nos2, nos1)
	fmt.Println(nos1, nos2)
	nos2[0] = 100
	fmt.Println(nos1, nos2)

}
