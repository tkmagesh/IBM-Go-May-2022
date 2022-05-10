package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter the name :")
	fmt.Scanln(&name)
	fmt.Printf("Hi %s, Have a nice day!\n", name)

	var n1, n2 int
	fmt.Println("Enter two numbers")
	//fmt.Scanln(&n1, &n2)
	fmt.Scanf("%d;%d\n", &n1, &n2)
	fmt.Println(n1, n2)
}
