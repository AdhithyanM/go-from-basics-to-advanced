package main

import "fmt"

func main() {
    // Array declaration
    var arr [5]int
    fmt.Println("Empty array:", arr)
    fmt.Printf("Array length: %d\n", len(arr))
    fmt.Printf("Array capacity: %d\n", cap(arr))

    // Array initialization
    numbers := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Initialized array:", numbers)

    // Array with size inference
    fruits := [...]string{"apple", "banana", "orange"}
    fmt.Println("Array with inferred size:", fruits)

    // Array iteration
    fmt.Println("\nIterating over array:")
    for i, fruit := range fruits {
        fmt.Printf("Index %d: %s\n", i, fruit)
    }

    // Slice declaration
    var slice []int
    fmt.Println("\nEmpty slice:", slice)
    fmt.Printf("Slice length: %d\n", len(slice))
    fmt.Printf("Slice capacity: %d\n", cap(slice))

    // Slice initialization
    numbersSlice := []int{1, 2, 3, 4, 5}
    fmt.Println("Initialized slice:", numbersSlice)

    // Creating slice from array
    arrSlice := numbers[1:4]
    fmt.Println("Slice from array:", arrSlice)

    // Slice operations
    fmt.Println("\nSlice operations:")
    
    // Append
    slice = append(slice, 1, 2, 3)
    fmt.Println("After append:", slice)

    // Append slice to slice
    slice = append(slice, numbersSlice...)
    fmt.Println("After appending another slice:", slice)

    // Slice capacity and length
    fmt.Printf("Slice length: %d\n", len(slice))
    fmt.Printf("Slice capacity: %d\n", cap(slice))

    // Make slice with specific capacity
    newSlice := make([]int, 3, 6)
    fmt.Println("\nSlice with make:", newSlice)
    fmt.Printf("Length: %d, Capacity: %d\n", len(newSlice), cap(newSlice))

    // Slice manipulation
    fmt.Println("\nSlice manipulation:")
    
    // Copy slice
    copySlice := make([]int, len(numbersSlice))
    copy(copySlice, numbersSlice)
    fmt.Println("Copied slice:", copySlice)

    // Sub-slice
    subSlice := numbersSlice[1:3]
    fmt.Println("Sub-slice:", subSlice)

    // Slice with full capacity
    fullSlice := numbersSlice[:]
    fmt.Println("Full slice:", fullSlice)

    // 2D slice
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("\n2D slice (matrix):")
    for _, row := range matrix {
        fmt.Println(row)
    }
} 