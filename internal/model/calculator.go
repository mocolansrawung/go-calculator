package model

import "fmt"

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

func (c *Calculator) Multiplication() int {
	return c.NumberOne + c.NumberTwo
}

func (c *Calculator) PrintInfo() {
	fmt.Printf("number one: %v\n", c.NumberOne)
	fmt.Printf("number two: %v\n", c.NumberTwo)
}
