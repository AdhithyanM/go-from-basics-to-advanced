package main

import "fmt"

func main() {
    // Type inference with := operator
    name := "John"        // string
    age := 25            // int
    height := 1.75       // float64
    isStudent := true    // bool
    score := 95.5        // float64
    grade := 'A'         // rune (int32)

    // Print values and their types
    fmt.Printf("name: %v (Type: %T)\n", name, name)
    fmt.Printf("age: %v (Type: %T)\n", age, age)
    fmt.Printf("height: %v (Type: %T)\n", height, height)
    fmt.Printf("isStudent: %v (Type: %T)\n", isStudent, isStudent)
    fmt.Printf("score: %v (Type: %T)\n", score, score)
    fmt.Printf("grade: %v (Type: %T)\n", grade, grade)

    // Type inference with var keyword
    var (
        city = "New York"    // string
        temp = 23.5         // float64
        year = 2024         // int
    )

    fmt.Println("\nUsing var keyword:")
    fmt.Printf("city: %v (Type: %T)\n", city, city)
    fmt.Printf("temp: %v (Type: %T)\n", temp, temp)
    fmt.Printf("year: %v (Type: %T)\n", year, year)

    // Type inference in function calls
    fmt.Println("\nType inference in function calls:")
    printType("Hello")      // string
    printType(42)          // int
    printType(3.14)        // float64
    printType(true)        // bool
}

// Generic function to print value and its type
func printType(v interface{}) {
    fmt.Printf("Value: %v (Type: %T)\n", v, v)
} 