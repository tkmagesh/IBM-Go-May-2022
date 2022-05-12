package main

import (
	"fmt"
	"sort"
)

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

type Products []Product

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result = fmt.Sprintf("%s%s\n", result, product)
	}
	return result
}

/*
func (products Products) FilterCostlyProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

//Sort

//sort.Interface interface implementation
func (products Products) Len() int {
	return len(products)
}

//default comparison (by Id)
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//Sort By Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//Sort By Cost
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//Sort By Cost
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//Sort By Cost
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	default:
		sort.Sort(products)

	}
}

func (products Products) SortSlice(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	default:
		products.SortSlice("Id")

	}
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
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("Index of stove = ", products.IndexOf(stove))
	fmt.Println("Does stove included in the collection ? :", products.Includes(stove))

	//Filter
	fmt.Println("Filter")
	fmt.Println("Costly Products")
	//costlyProducts := products.FilterCostlyProducts()
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Stationary Products")
	//stationaryProducts := products.FilterStationaryProducts()
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts)

	fmt.Println("Any")
	fmt.Println("Are there any costly product ?:", products.Any(costlyProductPredicate))

	fmt.Println("All")
	fmt.Println("Are all the products costly products ? :", products.All(costlyProductPredicate))

	fmt.Printf("\nSort\n")
	/*
		fmt.Println("Default Sort")
		sort.Sort(products)
		fmt.Println(products)

		fmt.Println("By Cost")
		sort.Sort(ByCost{products})
		fmt.Println(products)

		fmt.Println("By Name")
		sort.Sort(ByName{products})
		fmt.Println(products)

		fmt.Println("By Units")
		sort.Sort(ByUnits{products})
		fmt.Println(products)

		fmt.Println("By Category")
		sort.Sort(ByCategory{products})
		fmt.Println(products)
	*/

	fmt.Println("Default Sort")
	//products.Sort("")
	products.SortSlice("")
	fmt.Println(products)

	fmt.Println("By Cost")
	//products.Sort("Cost")
	products.SortSlice("Cost")
	fmt.Println(products)

	fmt.Println("By Name")
	//products.Sort("Name")
	products.SortSlice("Name")
	fmt.Println(products)

	fmt.Println("By Units")
	//products.Sort("Units")
	products.SortSlice("Units")
	fmt.Println(products)

	fmt.Println("By Category")
	//products.Sort("Category")
	products.SortSlice("Category")
	fmt.Println(products)
}
