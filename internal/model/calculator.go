package model

type Calculator struct {
	NumberOne int
	NumberTwo int
}

func (c *Calculator) Subtract() int {
	return c.NumberOne - c.NumberTwo
}
