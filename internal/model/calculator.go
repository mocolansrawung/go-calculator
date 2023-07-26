package model

import (
	"fmt"
)

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

func (c *Calculator) Add() int {
	return c.NumberOne + c.NumberTwo

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
	return c.firstNumber + c.secondNumber
}

func (c *Calculator) PrintInfo() {
	fmt.Printf("First Number: %v\n", c.firstNumber)
	fmt.Printf("Second Number: %v\n", c.secondNumber)
}
