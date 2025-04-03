package main

import "fmt"

func main() {
    // Numeric type conversions
    var (
        intNum    int     = 42
        floatNum  float64 = 3.14
        byteNum   byte    = 65
    )

    // Converting between numeric types
    fmt.Println("Numeric type conversions:")
    fmt.Printf("int to float64: %v (Type: %T)\n", float64(intNum), float64(intNum))
    fmt.Printf("float64 to int: %v (Type: %T)\n", int(floatNum), int(floatNum))
    fmt.Printf("byte to int: %v (Type: %T)\n", int(byteNum), int(byteNum))
    fmt.Printf("int to byte: %v (Type: %T)\n", byte(intNum), byte(intNum))

    // String conversions
    fmt.Println("\nString conversions:")
    // Converting byte to string (ASCII)
    fmt.Printf("byte to string: %v (Type: %T)\n", string(byteNum), string(byteNum))
    
    // Converting rune to string (Unicode)
    runeVal := 'A'
    fmt.Printf("rune to string: %v (Type: %T)\n", string(runeVal), string(runeVal))

    // Converting string to byte slice
    str := "Hello"
    byteSlice := []byte(str)
    fmt.Printf("string to []byte: %v (Type: %T)\n", byteSlice, byteSlice)

    // Converting byte slice to string
    fmt.Printf("[]byte to string: %v (Type: %T)\n", string(byteSlice), string(byteSlice))

    // Type conversion with custom types
    type Celsius float64
    type Fahrenheit float64

    var (
        c Celsius    = 100
        f Fahrenheit = 212
    )

    fmt.Println("\nCustom type conversions:")
    // Converting between custom types
    fmt.Printf("Celsius to Fahrenheit: %v°F (Type: %T)\n", 
        Fahrenheit(c*9/5 + 32), Fahrenheit(c*9/5 + 32))
    fmt.Printf("Fahrenheit to Celsius: %v°C (Type: %T)\n", 
        Celsius((f - 32) * 5 / 9), Celsius((f - 32) * 5 / 9))
} 