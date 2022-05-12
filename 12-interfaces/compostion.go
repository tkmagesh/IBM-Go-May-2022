package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}

//fmt.Stringer interface implementation
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) ApplyDiscount(disountPercent float32) {
	product.Cost = product.Cost * ((100 - disountPercent) / 100)
}

//fmt.Stringer interface implementation
//overriding the Format() method of the base type (Product)
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
}

func main() {

	grapes := PerishableProduct{
		Product{200, "Grapes", 20, 40, "Food"},
		"2 Days",
	}

	fmt.Println(grapes)
	grapes.ApplyDiscount(10)
	fmt.Println(grapes)
}
