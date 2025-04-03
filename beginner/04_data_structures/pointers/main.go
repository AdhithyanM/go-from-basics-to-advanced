package main

import "fmt"

// Function demonstrating basic pointer operations
func pointerBasics() {
    // Declaring a pointer
    var ptr *int
    fmt.Println("Zero value of pointer:", ptr) // nil

    // Creating a variable and getting its address
    x := 42
    ptr = &x
    fmt.Printf("Address of x: %p\n", ptr)
    fmt.Printf("Value at address: %d\n", *ptr)

    // Modifying value through pointer
    *ptr = 100
    fmt.Printf("New value of x: %d\n", x)
}

// Function demonstrating pointer to struct
func structPointers() {
    type Person struct {
        Name string
        Age  int
    }

    // Creating a struct and getting its pointer
    person := Person{Name: "John", Age: 25}
    personPtr := &person

    // Accessing struct fields through pointer
    fmt.Printf("Name: %s, Age: %d\n", personPtr.Name, personPtr.Age)
    // Same as: fmt.Printf("Name: %s, Age: %d\n", (*personPtr).Name, (*personPtr).Age)

    // Modifying struct through pointer
    personPtr.Age = 26
    fmt.Printf("Updated age: %d\n", person.Age)
}

// Function demonstrating pointer to array
func arrayPointers() {
    // Creating an array and getting its pointer
    arr := [3]int{1, 2, 3}
    arrPtr := &arr

    // Accessing array elements through pointer
    fmt.Printf("First element: %d\n", arrPtr[0])
    // Same as: fmt.Printf("First element: %d\n", (*arrPtr)[0])

    // Modifying array through pointer
    arrPtr[1] = 20
    fmt.Printf("Modified array: %v\n", arr)
}

// Function demonstrating pointer to slice
func slicePointers() {
    // Creating a slice
    slice := []int{1, 2, 3}
    slicePtr := &slice

    // Modifying slice through pointer
    (*slicePtr)[0] = 10
    fmt.Printf("Modified slice: %v\n", slice)

    // Appending to slice through pointer
    *slicePtr = append(*slicePtr, 4)
    fmt.Printf("Appended slice: %v\n", slice)
}

// Function demonstrating pointer to map
func mapPointers() {
    // Creating a map
    m := map[string]int{"a": 1, "b": 2}
    mapPtr := &m

    // Modifying map through pointer
    (*mapPtr)["c"] = 3
    fmt.Printf("Modified map: %v\n", m)
}

// Function demonstrating pointer to function
func functionPointers() {
    // Function type
    type MathFunc func(int, int) int

    // Function variable
    var add MathFunc = func(a, b int) int {
        return a + b
    }

    // Pointer to function
    funcPtr := &add

    // Calling function through pointer
    result := (*funcPtr)(5, 3)
    fmt.Printf("Result of function call: %d\n", result)
}

// Interface type for interface pointer example
type Stringer interface {
    String() string
}

// Struct implementing Stringer
type StringPerson struct {
    Name string
}

func (p StringPerson) String() string {
    return p.Name
}

// Function demonstrating pointer to interface
func interfacePointers() {
    // Creating interface variable and pointer
    var s Stringer = StringPerson{Name: "John"}
    sPtr := &s

    // Using interface through pointer
    fmt.Printf("String value: %s\n", (*sPtr).String())
}

func main() {
    fmt.Println("=== Basic Pointer Operations ===")
    pointerBasics()

    fmt.Println("\n=== Struct Pointers ===")
    structPointers()

    fmt.Println("\n=== Array Pointers ===")
    arrayPointers()

    fmt.Println("\n=== Slice Pointers ===")
    slicePointers()

    fmt.Println("\n=== Map Pointers ===")
    mapPointers()

    fmt.Println("\n=== Function Pointers ===")
    functionPointers()

    fmt.Println("\n=== Interface Pointers ===")
    interfacePointers()
} 