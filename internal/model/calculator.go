package model

import "fmt"

type Calculator struct {
	firstNumber  int
	secondNumber int
}

func NewCalculator(firstNum, secondNum int) *Calculator {
	return &Calculator{firstNumber: firstNum, secondNumber: secondNum}
}

func (c *Calculator) Subtract() int {
	return c.firstNumber - c.secondNumber
}
func (c *Calculator) Addition() int {
	return c.firstNumber + c.secondNumber
}

func (c *Calculator) Multiply() int {
	return c.firstNumber + c.secondNumber
}

func (c *Calculator) PrintInfo() {
	fmt.Printf("First Number: %v\n", c.firstNumber)
	fmt.Printf("Second Number: %v\n", c.secondNumber)
}
