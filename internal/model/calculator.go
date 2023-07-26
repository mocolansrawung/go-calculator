package model

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
