package main

import "fmt"

func main() {
    // Map declaration
    var m map[string]int
    fmt.Println("Empty map:", m)
    fmt.Printf("Map length: %d\n", len(m))

    // Map initialization
    scores := map[string]int{
        "John":  95,
        "Alice": 88,
        "Bob":   92,
    }
    fmt.Println("Initialized map:", scores)

    // Map operations
    fmt.Println("\nMap operations:")

    // Add/Update
    scores["Charlie"] = 85
    fmt.Println("After adding Charlie:", scores)

    // Get value
    if score, exists := scores["John"]; exists {
        fmt.Printf("John's score: %d\n", score)
    }

    // Delete
    delete(scores, "Bob")
    fmt.Println("After deleting Bob:", scores)

    // Map iteration
    fmt.Println("\nIterating over map:")
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }

    // Map with structs
    type Person struct {
        Age  int
        City string
    }

    people := map[string]Person{
        "John": {Age: 25, City: "New York"},
        "Alice": {Age: 30, City: "London"},
    }
    fmt.Println("\nMap with structs:", people)

    // Nested maps
    nested := map[string]map[string]int{
        "A": {"x": 1, "y": 2},
        "B": {"x": 3, "y": 4},
    }
    fmt.Println("\nNested map:", nested)

    // Map with slice values
    sliceMap := map[string][]int{
        "even": {2, 4, 6, 8},
        "odd":  {1, 3, 5, 7},
    }
    fmt.Println("\nMap with slice values:", sliceMap)

    // Map with function values
    operations := map[string]func(int, int) int{
        "add": func(a, b int) int { return a + b },
        "sub": func(a, b int) int { return a - b },
        "mul": func(a, b int) int { return a * b },
    }
    fmt.Println("\nMap with function values:")
    fmt.Printf("Add(5, 3) = %d\n", operations["add"](5, 3))
    fmt.Printf("Sub(5, 3) = %d\n", operations["sub"](5, 3))
    fmt.Printf("Mul(5, 3) = %d\n", operations["mul"](5, 3))

    // Map with interface{} values
    mixed := map[string]interface{}{
        "name":   "John",
        "age":    25,
        "scores": []int{85, 92, 88},
    }
    fmt.Println("\nMap with interface{} values:", mixed)
} 