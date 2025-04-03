package main

import "fmt"

// Basic function with no parameters and no return value
func sayHello() {
    fmt.Println("Hello, Go!")
}

// Function with parameters
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Function with named return values
func rectangle(width, height float64) (area float64, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return
}

// Function with multiple parameters of same type
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Function that takes a function as parameter
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    // Calling basic function
    sayHello()

    // Calling function with parameters
    greet("John")

    // Using function with return value
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)

    // Using function with multiple return values
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", quotient)
    }

    // Using function with named return values
    area, perimeter := rectangle(5, 3)
    fmt.Printf("Rectangle: Area = %.2f, Perimeter = %.2f\n", area, perimeter)

    // Using variadic function
    total := sum(1, 2, 3, 4, 5)
    fmt.Printf("Sum of numbers: %d\n", total)

    // Using function that takes another function
    multiply := func(a, b int) int {
        return a * b
    }
    result = applyOperation(4, 2, multiply)
    fmt.Printf("4 * 2 = %d\n", result)
} 