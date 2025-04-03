package main

import "fmt"

func main() {
    // Zero values for basic types
    var (
        // Numeric types
        intNum    int
        floatNum  float64
        complexNum complex128

        // String type
        str string

        // Boolean type
        boolVal bool

        // Rune type (alias for int32)
        runeVal rune

        // Byte type (alias for uint8)
        byteVal byte
    )

    fmt.Println("Zero values for basic types:")
    fmt.Printf("int: %v (Type: %T)\n", intNum, intNum)
    fmt.Printf("float64: %v (Type: %T)\n", floatNum, floatNum)
    fmt.Printf("complex128: %v (Type: %T)\n", complexNum, complexNum)
    fmt.Printf("string: %q (Type: %T)\n", str, str)
    fmt.Printf("bool: %v (Type: %T)\n", boolVal, boolVal)
    fmt.Printf("rune: %v (Type: %T)\n", runeVal, runeVal)
    fmt.Printf("byte: %v (Type: %T)\n", byteVal, byteVal)

    // Zero values for composite types
    var (
        // Array
        arr [3]int

        // Slice
        slice []int

        // Map
        m map[string]int

        // Channel
        ch chan int

        // Pointer
        ptr *int

        // Function
        fn func()
    )

    fmt.Println("\nZero values for composite types:")
    fmt.Printf("array: %v (Type: %T)\n", arr, arr)
    fmt.Printf("slice: %v (Type: %T)\n", slice, slice)
    fmt.Printf("map: %v (Type: %T)\n", m, m)
    fmt.Printf("channel: %v (Type: %T)\n", ch, ch)
    fmt.Printf("pointer: %v (Type: %T)\n", ptr, ptr)
    fmt.Printf("function: %v (Type: %T)\n", fn, fn)

    // Zero values for custom types
    type Person struct {
        Name string
        Age  int
    }

    var person Person
    fmt.Println("\nZero values for custom type:")
    fmt.Printf("person: %+v (Type: %T)\n", person, person)
} 