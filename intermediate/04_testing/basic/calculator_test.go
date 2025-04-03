package calculator

import "testing"

func TestAdd(t *testing.T) {
    calc := &Calculator{}
    result := calc.Add(2, 3)
    expected := 5.0
    
    if result != expected {
        t.Errorf("Add(2, 3) = %f; expected %f", result, expected)
    }
    
    if calc.GetLastResult() != expected {
        t.Errorf("GetLastResult() = %f; expected %f", calc.GetLastResult(), expected)
    }
}

func TestSubtract(t *testing.T) {
    calc := &Calculator{}
    result := calc.Subtract(5, 3)
    expected := 2.0
    
    if result != expected {
        t.Errorf("Subtract(5, 3) = %f; expected %f", result, expected)
    }
}

func TestMultiply(t *testing.T) {
    calc := &Calculator{}
    result := calc.Multiply(2, 3)
    expected := 6.0
    
    if result != expected {
        t.Errorf("Multiply(2, 3) = %f; expected %f", result, expected)
    }
}

func TestDivide(t *testing.T) {
    calc := &Calculator{}
    
    // Test valid division
    result, err := calc.Divide(6, 2)
    if err != nil {
        t.Errorf("Divide(6, 2) returned unexpected error: %v", err)
    }
    expected := 3.0
    if result != expected {
        t.Errorf("Divide(6, 2) = %f; expected %f", result, expected)
    }
    
    // Test division by zero
    _, err = calc.Divide(1, 0)
    if err == nil {
        t.Error("Divide(1, 0) should return error for division by zero")
    }
} 