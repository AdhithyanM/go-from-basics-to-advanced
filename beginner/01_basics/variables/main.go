package main

import "fmt"

func main() {
    // Variable declarations
    var name string = "John"
    var age int = 25
    var height float64 = 1.75
    var isStudent bool = true

    // Short variable declaration (:=)
    city := "New York"
    temperature := 23.5

    // Multiple variable declaration
    var (
        firstName = "Jane"
        lastName  = "Doe"
        email     = "jane@example.com"
    )

    // Printing variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("City:", city)
    fmt.Println("Temperature:", temperature)
    fmt.Println("Full Name:", firstName, lastName)
    fmt.Println("Email:", email)

    // Demonstrating zero values
    var (
        defaultInt    int
        defaultFloat  float64
        defaultString string
        defaultBool   bool
    )

    fmt.Println("\nZero Values:")
    fmt.Printf("Default Int: %v\n", defaultInt)
    fmt.Printf("Default Float: %v\n", defaultFloat)
    fmt.Printf("Default String: %q\n", defaultString)
    fmt.Printf("Default Bool: %v\n", defaultBool)
}