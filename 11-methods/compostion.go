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

func main() {

	grapes := PerishableProduct{
		Product{200, "Grapes", 20, 40, "Food"},
		"2 Days",
	}
	//fmt.Println(grapes.Product.Format())

	fmt.Println(grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) ApplyDiscount(disountPercent float32) {
	//(*product).Cost = (*product).Cost * ((100 - disountPercent) / 100)
	product.Cost = product.Cost * ((100 - disountPercent) / 100)
}

//overriding the Format() method of the base type (Product)
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

/*

//the below function is redundant when "ApplyDiscount" is implemented as a method
func (pp *PerishableProduct) ApplyDiscountPP(disountPercent float32) {
	//(*product).Cost = (*product).Cost * ((100 - disountPercent) / 100)
	pp.Product.Cost = pp.Product.Cost * ((100 - disountPercent) / 100)
}
*/
