package main

import "fmt"

// Constants can be declared at package level
const (
    // Numeric constants
    Pi       = 3.14159
    MaxValue = 1000
    
    // String constants
    Greeting = "Hello, Go!"
    
    // Boolean constants
    IsEnabled = true
)

func main() {
    // Using package-level constants
    fmt.Println("Pi:", Pi)
    fmt.Println("Max Value:", MaxValue)
    fmt.Println("Greeting:", Greeting)
    fmt.Println("Is Enabled:", IsEnabled)

    // Constants can be declared inside functions too
    const (
        // Typed constants
        Age     int     = 25
        Height  float64 = 1.75
        Name    string  = "John"
        
        // Untyped constants (type will be inferred)
        Temperature = 23.5
        IsStudent   = true
    )

    fmt.Println("\nFunction-level constants:")
    fmt.Printf("Age: %v (Type: %T)\n", Age, Age)
    fmt.Printf("Height: %v (Type: %T)\n", Height, Height)
    fmt.Printf("Name: %v (Type: %T)\n", Name, Name)
    fmt.Printf("Temperature: %v (Type: %T)\n", Temperature, Temperature)
    fmt.Printf("Is Student: %v (Type: %T)\n", IsStudent, IsStudent)

    // Constants can be used in expressions
    const (
        Celsius    = 100
        Fahrenheit = (Celsius * 9 / 5) + 32
    )

    fmt.Printf("\nTemperature conversion:\n")
    fmt.Printf("%v°C = %v°F\n", Celsius, Fahrenheit)
} 