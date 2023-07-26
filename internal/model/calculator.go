package model

type Calculator struct {
	numberOne int
	numberTwo int
}

func (c *Calculator) Subtract() int {
	return c.numberOne - c.numberTwo
}
