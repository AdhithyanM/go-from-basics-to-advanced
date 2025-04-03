package main

import "fmt"

func main() {
    // Basic if statement
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult")
    }

    // if-else statement
    score := 75
    if score >= 60 {
        fmt.Println("You passed!")
    } else {
        fmt.Println("You failed!")
    }

    // if-else if-else statement
    grade := 85
    if grade >= 90 {
        fmt.Println("Grade: A")
    } else if grade >= 80 {
        fmt.Println("Grade: B")
    } else if grade >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }

    // if with initialization
    if num := 42; num > 40 {
        fmt.Printf("%d is greater than 40\n", num)
    }
    // Note: num is not accessible here

    // Nested if statements
    isStudent := true
    hasID := true
    if isStudent {
        if hasID {
            fmt.Println("Student has valid ID")
        } else {
            fmt.Println("Student needs ID")
        }
    } else {
        fmt.Println("Not a student")
    }

    // Multiple conditions in if
    age = 25
    hasLicense := true
    if age >= 18 && hasLicense {
        fmt.Println("Can drive")
    }

    // Using logical operators
    isWeekend := true
    isHoliday := false
    if isWeekend || isHoliday {
        fmt.Println("No work today!")
    }

    // Complex conditions
    temperature := 25
    humidity := 60
    if temperature > 30 && humidity > 70 {
        fmt.Println("It's too hot and humid!")
    } else if temperature < 10 || humidity < 30 {
        fmt.Println("It's too cold or dry!")
    } else {
        fmt.Println("The weather is comfortable")
    }
} 