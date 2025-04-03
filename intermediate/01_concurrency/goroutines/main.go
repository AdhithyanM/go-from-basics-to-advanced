package main

import (
	"fmt"
	"time"
)

// Basic goroutine function
func sayHello() {
    fmt.Println("Hello from goroutine!")
}

// Function with parameter
func printNumber(num int) {
    fmt.Printf("Number: %d\n", num)
}

// Function demonstrating goroutine lifecycle
func lifecycleExample() {
    fmt.Println("Starting lifecycle example")
    
    // Goroutine that completes before main
    go func() {
        fmt.Println("Quick goroutine")
    }()

    // Goroutine that might not complete before main
    go func() {
        fmt.Println("Slow goroutine")
        time.Sleep(100 * time.Millisecond)
    }()

    // Wait to see the output
    time.Sleep(50 * time.Millisecond)
}

// Function demonstrating multiple goroutines
func multipleGoroutines() {
    fmt.Println("\nMultiple goroutines example:")
    
    // Launch multiple goroutines
    for i := 0; i < 5; i++ {
        go func(num int) {
            fmt.Printf("Goroutine %d started\n", num)
            time.Sleep(100 * time.Millisecond)
            fmt.Printf("Goroutine %d finished\n", num)
        }(i)
    }

    // Wait to see the output
    time.Sleep(200 * time.Millisecond)
}

// Function demonstrating goroutine with closure
func closureExample() {
    fmt.Println("\nClosure example:")
    
    message := "Hello from closure"
    go func() {
        fmt.Println(message)
    }()
    
    time.Sleep(50 * time.Millisecond)
}

// Function demonstrating goroutine with loop variable
func loopVariableExample() {
    fmt.Println("\nLoop variable example:")
    
    // Incorrect way (common mistake)
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Printf("Incorrect: %d\n", i)
        }()
    }
    
    time.Sleep(100 * time.Millisecond)
    
    // Correct way
    for i := 0; i < 3; i++ {
        go func(num int) {
            fmt.Printf("Correct: %d\n", num)
        }(i)
    }
    
    time.Sleep(100 * time.Millisecond)
}

// Function demonstrating goroutine with panic handling
func panicExample() {
    fmt.Println("\nPanic handling example:")
    
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Recovered from panic: %v\n", r)
            }
        }()
        
        panic("Something went wrong!")
    }()
    
    time.Sleep(100 * time.Millisecond)
}

func main() {
    // Basic goroutine
    fmt.Println("Basic goroutine example:")
    go sayHello()
    time.Sleep(50 * time.Millisecond)

    // Goroutine with parameter
    fmt.Println("\nGoroutine with parameter:")
    go printNumber(42)
    time.Sleep(50 * time.Millisecond)

    // Goroutine lifecycle
    lifecycleExample()

    // Multiple goroutines
    multipleGoroutines()

    // Closure example
    closureExample()

    // Loop variable example
    loopVariableExample()

    // Panic handling
    panicExample()
} 