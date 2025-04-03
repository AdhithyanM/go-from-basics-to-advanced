package calculator

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 0, 0},
        {"decimal numbers", 1.5, 2.5, 4.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calc := &Calculator{}
            result := calc.Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%f, %f) = %f; expected %f", 
                    tt.a, tt.b, result, tt.expected)
            }
            if calc.GetLastResult() != tt.expected {
                t.Errorf("GetLastResult() = %f; expected %f", 
                    calc.GetLastResult(), tt.expected)
            }
        })
    }
}

func TestSubtract(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"positive numbers", 5, 3, 2},
        {"negative numbers", -2, -3, 1},
        {"zero", 0, 0, 0},
        {"decimal numbers", 3.5, 1.5, 2.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calc := &Calculator{}
            result := calc.Subtract(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Subtract(%f, %f) = %f; expected %f", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"positive numbers", 2, 3, 6},
        {"negative numbers", -2, -3, 6},
        {"zero", 0, 5, 0},
        {"decimal numbers", 1.5, 2, 3.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calc := &Calculator{}
            result := calc.Multiply(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Multiply(%f, %f) = %f; expected %f", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestDivide(t *testing.T) {
    tests := []struct {
        name        string
        a, b        float64
        expected    float64
        expectError bool
    }{
        {"valid division", 6, 2, 3, false},
        {"division by zero", 1, 0, 0, true},
        {"decimal division", 5, 2, 2.5, false},
        {"negative division", -6, 2, -3, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calc := &Calculator{}
            result, err := calc.Divide(tt.a, tt.b)
            
            // Check error expectation
            if tt.expectError && err == nil {
                t.Error("expected error but got none")
            }
            if !tt.expectError && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            
            // Check result if no error expected
            if !tt.expectError && result != tt.expected {
                t.Errorf("Divide(%f, %f) = %f; expected %f", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
} 