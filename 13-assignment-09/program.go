package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
write the apis for the following APIs

        IndexOf => return the index of the given product
            ex: IndexOf(product) => returns the index of the given product

        Includes => return true if the given product is present in the collection else return false
            ex: Includes(product) => returns true if the given product is present in the collection

        Filter => return a new collection of products that satisfy the given condition
            use cases:
                1. filter all costly products (cost > 1000)
                2. filter all stationary products (category = "Stationary")
                3. filter all costly stationary products
				etc

        Any => return true if any of the product in the collections satifies the given criteria
            use cases:
                1. are there any costly product (cost > 1000)?
                2. are there any stationary product (category = "Stationary")?
                3. are there any costly stationary product?
				etc

        All => return true if all the products in the collections satifies the given criteria
            use cases:
                1. are all the products costly products (cost > 1000)?
                2. are all the products stationary products (category = "Stationary")?
                3. are all the products costly stationary products?
				etc

        Sort => Sort the products collection by any attribute
			IMPORTANT : Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order

*/

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
}