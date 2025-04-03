package main

import "fmt"

// Defer example
func deferExample() {
    fmt.Println("Starting defer example")
    defer fmt.Println("This will be printed last")
    defer fmt.Println("This will be printed second to last")
    fmt.Println("This will be printed first")
}

// Defer with function
func deferWithFunction() {
    fmt.Println("Starting function")
    defer func() {
        fmt.Println("Deferred function")
    }()
    fmt.Println("Ending function")
}

// Panic and recover example
func panicExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    fmt.Println("About to panic")
    panic("Something went wrong!")
    fmt.Println("This won't be printed")
}

// Function composition
func addOne(x int) int {
    return x + 1
}

func multiplyByTwo(x int) int {
    return x * 2
}

func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

// Function as value
type MathFunc func(int) int

func applyFunction(x int, f MathFunc) int {
    return f(x)
}

// Defer with file operations
func fileOperations() {
    fmt.Println("Opening file")
    defer fmt.Println("Closing file")
    fmt.Println("Writing to file")
}

// Multiple defer statements
func multipleDefer() {
    for i := 0; i < 3; i++ {
        defer fmt.Printf("Deferred %d\n", i)
    }
    fmt.Println("Function body")
}

func main() {
    // Defer examples
    fmt.Println("=== Defer Examples ===")
    deferExample()
    deferWithFunction()
    fileOperations()
    multipleDefer()

    // Panic and recover
    fmt.Println("\n=== Panic and Recover Example ===")
    panicExample()

    // Function composition
    fmt.Println("\n=== Function Composition ===")
    composed := compose(addOne, multiplyByTwo)
    result := composed(5)
    fmt.Printf("Composed function result: %d\n", result)

    // Function as value
    fmt.Println("\n=== Function as Value ===")
    square := func(x int) int {
        return x * x
    }
    result = applyFunction(4, square)
    fmt.Printf("Square of 4: %d\n", result)

    // Defer with return values
    fmt.Println("\n=== Defer with Return Values ===")
    result = func() int {
        defer fmt.Println("Deferred in return function")
        return 42
    }()
    fmt.Printf("Return value: %d\n", result)
} 