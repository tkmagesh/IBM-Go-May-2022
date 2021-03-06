package main

import (
	"fmt"
	"modularity-demo/calc"
	calcUtils "modularity-demo/calc/utils"
	_ "modularity-demo/dummy"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Printf("OpCount = %d\n", calc.OpCount())
	fmt.Println("Is 21 an even number ? :", calcUtils.IsEven(21))
	color.Red("This text will be printed in red color")
	color.Blue("This text will be printed in blue color")
}
