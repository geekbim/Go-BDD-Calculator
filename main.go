package main

type Calculator struct {
	result int
}

func (c Calculator) Result() int {
	return c.result
}

func (c *Calculator) Add(num int) {
	c.result += num
}

func (c *Calculator) Subtract(num int) {
	c.result -= num
}

func (c *Calculator) Multiply(num int) {
	c.result *= num
}

func (c *Calculator) Divide(num int) {
	c.result /= num
}

func NewCalculator(num int) *Calculator {
	return &Calculator{num}
}
