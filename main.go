package main

import (
	"fmt"
	"go-calculator/internal/model"
)

func main() {
	calc := model.NewCalculator(1, 3)
	fmt.Println(calc.Subtract())
}
