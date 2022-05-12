package main

import "fmt"

func main() {
	var x any
	x = 100
	x = 10.98
	x = "Magesh"
	x = true
	fmt.Println(x)

	//type assertion
	var z any //in go 1.18
	//var z interface{} // below go 1.17

	//z = 100
	z = "Some string"
	if val, ok := z.(int); ok {
		fmt.Println(val + 200)
	} else {
		fmt.Println("z is not an int")
	}

	//z = 100
	//z = "some string"
	//z = float32(100.05)
	//z = true
	//z = struct{}{}
	//z = []int{3, 1, 4, 2, 5}
	z = int64(100)
	switch val := z.(type) {
	case int:
		fmt.Println("z is an int, z + 200 = ", val+200)
	case string:
		fmt.Println("z is a string, len(z) = ", len(val))
	case float32:
		fmt.Println("z is a float32, z + 200.9999 = ", val+200.9999)
	case bool:
		fmt.Println("z is a bool, !z  = ", !val)
	case struct{}:
		fmt.Println("z is an empty struct")
	case []int:
		fmt.Println("z is a []int, z = ", val)
	default:
		fmt.Println("Unknown type")
	}
}
