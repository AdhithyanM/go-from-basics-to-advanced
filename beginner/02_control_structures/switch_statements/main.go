package main

import "fmt"

func main() {
    // Basic switch statement
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("End of the week")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Midweek")
    }

    // Switch with expression
    score := 85
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80:
        fmt.Println("Grade: B")
    case score >= 70:
        fmt.Println("Grade: C")
    default:
        fmt.Println("Grade: F")
    }

    // Switch with initialization
    switch num := 42; {
    case num < 0:
        fmt.Println("Negative number")
    case num > 0:
        fmt.Println("Positive number")
    default:
        fmt.Println("Zero")
    }

    // Switch with multiple cases
    month := "January"
    switch month {
    case "December", "January", "February":
        fmt.Println("Winter")
    case "March", "April", "May":
        fmt.Println("Spring")
    case "June", "July", "August":
        fmt.Println("Summer")
    case "September", "October", "November":
        fmt.Println("Fall")
    }

    // Switch with fallthrough
    grade := "B"
    switch grade {
    case "A":
        fmt.Println("Excellent!")
        fallthrough
    case "B":
        fmt.Println("Good!")
        fallthrough
    case "C":
        fmt.Println("Fair")
    case "F":
        fmt.Println("Failed")
    }

    // Type switch
    var i interface{} = "hello"
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %v\n", v)
    case string:
        fmt.Printf("String: %v\n", v)
    case float64:
        fmt.Printf("Float: %v\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }

    // Switch with multiple expressions
    char := 'a'
    switch char {
    case 'a', 'e', 'i', 'o', 'u':
        fmt.Printf("%c is a vowel\n", char)
    case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm',
         'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
        fmt.Printf("%c is a consonant\n", char)
    default:
        fmt.Printf("%c is not a letter\n", char)
    }
} 