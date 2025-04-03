package main

import (
	"encoding/json"
	"fmt"
)

// Basic struct
type Person struct {
    Name string
    Age  int
}

// Struct with tags
type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"-"` // Private field
}

// Embedded struct
type Address struct {
    Street  string
    City    string
    Country string
}

type Contact struct {
    Phone   string
    Email   string
    Address Address // Embedded struct
}

// Struct with methods
type Rectangle struct {
    Width  float64
    Height float64
}

// Method on Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Struct with interface
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func main() {
    // Creating struct instances
    person := Person{
        Name: "John",
        Age:  25,
    }
    fmt.Println("Person:", person)

    // Struct with tags
    user := User{
        ID:       1,
        Username: "john_doe",
        Email:    "john@example.com",
        Password: "secret",
    }
    fmt.Println("User:", user)

    // JSON marshaling
    jsonData, _ := json.Marshal(user)
    fmt.Println("JSON:", string(jsonData))

    // Embedded struct
    contact := Contact{
        Phone: "123-456-7890",
        Email: "john@example.com",
        Address: Address{
            Street:  "123 Main St",
            City:    "New York",
            Country: "USA",
        },
    }
    fmt.Println("Contact:", contact)

    // Accessing embedded fields
    fmt.Printf("City: %s\n", contact.Address.City)

    // Struct with methods
    rect := Rectangle{Width: 5, Height: 3}
    fmt.Printf("Rectangle area: %.2f\n", rect.Area())

    // Using pointer receiver
    rect.Scale(2)
    fmt.Printf("Scaled rectangle area: %.2f\n", rect.Area())

    // Interface implementation
    shapes := []Shape{
        Rectangle{Width: 5, Height: 3},
        Circle{Radius: 2.5},
    }

    fmt.Println("\nAreas of shapes:")
    for _, shape := range shapes {
        fmt.Printf("Area: %.2f\n", shape.Area())
    }

    // Anonymous struct
    car := struct {
        Make  string
        Model string
        Year  int
    }{
        Make:  "Toyota",
        Model: "Camry",
        Year:  2020,
    }
    fmt.Println("\nAnonymous struct:", car)

    // Struct comparison
    p1 := Person{Name: "John", Age: 25}
    p2 := Person{Name: "John", Age: 25}
    fmt.Printf("\nStruct comparison: %v\n", p1 == p2)
} 