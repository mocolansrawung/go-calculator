package main

import (
	"fmt"
	"go-calculator/internal/model"
)

func main() {
	calc := model.NewCalculator(1, 3)

	fmt.Println(calc.Subtract())
  	fmt.Println(calc.Add())
	fmt.Println(calc.Divide())
	fmt.Println(calc.Multiply())

	calc.PrintInfo()
}
