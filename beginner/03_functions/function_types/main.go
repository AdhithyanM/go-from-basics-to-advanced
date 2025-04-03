package main

import "fmt"

// Variadic function
func printNumbers(numbers ...int) {
    fmt.Println("Numbers:", numbers)
    for i, num := range numbers {
        fmt.Printf("Index %d: %d\n", i, num)
    }
}

// Anonymous function
var square = func(x int) int {
    return x * x
}

// Closure example
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Method example
type Rectangle struct {
    width  float64
    height float64
}

// Method on Rectangle type
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}

// Function type
type Operation func(int, int) int

func main() {
    // Using variadic function
    fmt.Println("Variadic function example:")
    printNumbers(1, 2, 3, 4, 5)
    printNumbers() // Can be called with no arguments

    // Using anonymous function
    fmt.Println("\nAnonymous function example:")
    fmt.Printf("Square of 5: %d\n", square(5))

    // Using closure
    fmt.Println("\nClosure example:")
    counter := createCounter()
    fmt.Printf("Count: %d\n", counter())
    fmt.Printf("Count: %d\n", counter())
    fmt.Printf("Count: %d\n", counter())

    // Using methods
    fmt.Println("\nMethod example:")
    rect := Rectangle{width: 5, height: 3}
    fmt.Printf("Area: %.2f\n", rect.Area())
    
    // Using method with pointer receiver
    rect.Scale(2)
    fmt.Printf("Area after scaling: %.2f\n", rect.Area())

    // Using function type
    fmt.Println("\nFunction type example:")
    var add Operation = func(a, b int) int {
        return a + b
    }
    fmt.Printf("Add: %d\n", add(5, 3))

    // Anonymous function as a goroutine
    fmt.Println("\nAnonymous function as goroutine:")
    go func() {
        fmt.Println("Running in goroutine")
    }()

    // Anonymous function with parameters
    func(x int) {
        fmt.Printf("Anonymous function with parameter: %d\n", x)
    }(42)
} 