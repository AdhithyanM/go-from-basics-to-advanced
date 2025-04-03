package calculator

import "errors"

// Calculator represents a basic calculator
type Calculator struct {
    lastResult float64
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b float64) float64 {
    c.lastResult = a + b
    return c.lastResult
}

// Subtract returns the difference between two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
    c.lastResult = a - b
    return c.lastResult
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
    c.lastResult = a * b
    return c.lastResult
}

// Divide returns the quotient of two numbers
func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    c.lastResult = a / b
    return c.lastResult, nil
}

// GetLastResult returns the result of the last operation
func (c *Calculator) GetLastResult() float64 {
    return c.lastResult
} 