package main

import "fmt"

func main() {
    // Basic break
    fmt.Println("Basic break:")
    for i := 0; i < 5; i++ {
        if i == 3 {
            break
        }
        fmt.Printf("i: %d\n", i)
    }

    // Basic continue
    fmt.Println("\nBasic continue:")
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue
        }
        fmt.Printf("i: %d\n", i)
    }

    // Break in nested loops
    fmt.Println("\nBreak in nested loops:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break // Only breaks the inner loop
            }
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }

    // Labeled break
    fmt.Println("\nLabeled break:")
    outer:
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if i == 1 && j == 1 {
                    break outer // Breaks both loops
                }
                fmt.Printf("i: %d, j: %d\n", i, j)
            }
        }

    // Continue in nested loops
    fmt.Println("\nContinue in nested loops:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                continue // Only skips the current iteration of inner loop
            }
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }

    // Labeled continue
    fmt.Println("\nLabeled continue:")
    outerContinue:
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if i == 1 && j == 1 {
                    continue outerContinue // Skips to next iteration of outer loop
                }
                fmt.Printf("i: %d, j: %d\n", i, j)
            }
        }

    // Break in switch statement
    fmt.Println("\nBreak in switch statement:")
    switch {
    case true:
        fmt.Println("First case")
        break
        fmt.Println("This won't be printed")
    case false:
        fmt.Println("Second case")
    }

    // Continue in for range loop
    fmt.Println("\nContinue in for range loop:")
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        if num == 3 {
            continue
        }
        fmt.Printf("Number: %d\n", num)
    }
} 