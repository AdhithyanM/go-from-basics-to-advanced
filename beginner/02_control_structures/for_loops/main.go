package main

import "fmt"

func main() {
    // Basic for loop
    fmt.Println("Basic for loop:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Iteration %d\n", i)
    }

    // For loop as while
    fmt.Println("\nFor loop as while:")
    count := 0
    for count < 3 {
        fmt.Printf("Count is %d\n", count)
        count++
    }

    // Infinite loop with break
    fmt.Println("\nInfinite loop with break:")
    counter := 0
    for {
        if counter >= 3 {
            break
        }
        fmt.Printf("Counter is %d\n", counter)
        counter++
    }

    // Range-based for loop with slice
    fmt.Println("\nRange-based for loop with slice:")
    fruits := []string{"apple", "banana", "orange"}
    for index, value := range fruits {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }

    // Range-based for loop with map
    fmt.Println("\nRange-based for loop with map:")
    person := map[string]string{
        "name":  "John",
        "age":   "25",
        "city": "New York",
    }
    for key, value := range person {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }

    // Range-based for loop with string
    fmt.Println("\nRange-based for loop with string:")
    for index, char := range "Hello" {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }

    // Nested loops
    fmt.Println("\nNested loops:")
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }

    // For loop with multiple variables
    fmt.Println("\nFor loop with multiple variables:")
    for i, j := 0, 10; i < j; i, j = i+1, j-1 {
        fmt.Printf("i: %d, j: %d\n", i, j)
    }

    // Range-based for loop with channel
    fmt.Println("\nRange-based for loop with channel:")
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
} 