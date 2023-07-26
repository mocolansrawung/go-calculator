package main

import (
	"fmt"
	"go-calculator/internal/model"
)

func main() {
	calc := model.Calculator{NumberOne: 1, NumberTwo: 3}
	fmt.Println(calc.Subtract())

}
