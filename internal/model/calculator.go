package model

import (
	"fmt"
)

type Calculator struct {
	NumberOne int
	NumberTwo int
}

func NewCalculator(numOne, numTwo int) *Calculator {
	return &Calculator{NumberOne: numOne, NumberTwo: numTwo}
}

func (c *Calculator) Subtract() int {
	return c.NumberOne - c.NumberTwo
}
func (c *Calculator) Addition() int {
	return c.NumberOne + c.NumberTwo
}

func (c *Calculator) Divide() float64 {
	result := float64(c.NumberOne) / float64(c.NumberTwo)
	formattedResult := fmt.Sprintf("%.2f", result)

	var roundedResult float64
    _, err := fmt.Sscanf(formattedResult, "%f", &roundedResult)
    if err != nil {
        fmt.Println(err)
    }
	return roundedResult  
}


func (c *Calculator) Multiply() int {
	return c.NumberOne + c.NumberTwo
}

func (c *Calculator) PrintInfo() {
	fmt.Printf("number one: %v\n", c.NumberOne)
	fmt.Printf("number two: %v\n", c.NumberTwo)
}
