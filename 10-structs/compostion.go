package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type Dummy struct {
	Cost float32
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
	/*
		grapes := PerishableProduct{
			Product{200, "Grapes", 20, 40, "Food"},
			"2 Days",
		}
	*/
	/* grapes := PerishableProduct{
		Product: Product{200, "Grapes", 20, 40, "Food"},
		Expiry:  "2 Days",
	} */

	grapes := NewPerishableProduct(200, "Grapes", 20, 40, "Food", "2 Days")
	fmt.Println(grapes)
	fmt.Println(grapes.Product.Cost)
	fmt.Println(grapes.Cost)

	/*
		fmt.Println(Format(grapes.Product))
		ApplyDiscount(&(grapes.Product), 10)
		fmt.Println(Format(grapes.Product))
	*/

	fmt.Println(FormatPP(*grapes))
	ApplyDiscountPP(grapes, 10)
	fmt.Println(FormatPP(*grapes))
}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func ApplyDiscount(product *Product, disountPercent float32) {
	//(*product).Cost = (*product).Cost * ((100 - disountPercent) / 100)
	product.Cost = product.Cost * ((100 - disountPercent) / 100)
}

func FormatPP(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry = %q", Format(pp.Product), pp.Expiry)
}

func ApplyDiscountPP(pp *PerishableProduct, disountPercent float32) {
	//(*product).Cost = (*product).Cost * ((100 - disountPercent) / 100)
	pp.Product.Cost = pp.Product.Cost * ((100 - disountPercent) / 100)
}
