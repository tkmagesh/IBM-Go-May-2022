package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	/* Create a Product struct with the following fields
		id, name, cost, units, category

	   Write a "Format" function that will for the product object to a string for printing in the console
	   	=> Id = 100, Name = "Pen", Cost = 10, Units = 100, Category = "Stationary"

	   Write a "ApplyDiscount" function to apply a discount on the given product
	   ApplyDiscount(p, 10) => Cost reduced by 10%
	*/

	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}
	fmt.Println("Before applying 10% discount", pen.Format())
	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	fmt.Println("After applying 10% discount", pen.Format())

	penPtr := &pen
	fmt.Println(penPtr.Name)

	//productPtr := new(Product)
	//productPtr := &Product{}
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) ApplyDiscount(disountPercent float32) {
	//(*product).Cost = (*product).Cost * ((100 - disountPercent) / 100)
	product.Cost = product.Cost * ((100 - disountPercent) / 100)
}
