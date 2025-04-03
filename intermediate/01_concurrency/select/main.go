package main

import (
	"fmt"
	"time"
)

// Basic select
func basicSelect() {
    fmt.Println("Basic select:")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        ch1 <- "First channel"
    }()
    
    go func() {
        ch2 <- "Second channel"
    }()
    
    select {
    case msg1 := <-ch1:
        fmt.Printf("Received from ch1: %s\n", msg1)
    case msg2 := <-ch2:
        fmt.Printf("Received from ch2: %s\n", msg2)
    }
}

// Select with default case
func selectWithDefault() {
    fmt.Println("\nSelect with default case:")
    
    ch := make(chan string)
    
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    default:
        fmt.Println("No message available")
    }
}

// Select with timeout
func selectWithTimeout() {
    fmt.Println("\nSelect with timeout:")
    
    ch := make(chan string)
    
    // Try to receive with timeout
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout waiting for message")
    }
}

// Select with multiple cases
func selectMultipleCases() {
    fmt.Println("\nSelect with multiple cases:")
    
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    
    go func() {
        ch1 <- 1
    }()
    
    go func() {
        ch2 <- 2
    }()
    
    go func() {
        ch3 <- 3
    }()
    
    for i := 0; i < 3; i++ {
        select {
        case val1 := <-ch1:
            fmt.Printf("Received from ch1: %d\n", val1)
        case val2 := <-ch2:
            fmt.Printf("Received from ch2: %d\n", val2)
        case val3 := <-ch3:
            fmt.Printf("Received from ch3: %d\n", val3)
        }
    }
}

// Select with send operations
func selectWithSend() {
    fmt.Println("\nSelect with send operations:")
    
    ch := make(chan int, 1)
    
    select {
    case ch <- 42:
        fmt.Println("Sent value 42")
    default:
        fmt.Println("Channel is full")
    }
    
    // Try to receive
    select {
    case val := <-ch:
        fmt.Printf("Received: %d\n", val)
    default:
        fmt.Println("No value available")
    }
}

// Select with multiple timeouts
func selectWithMultipleTimeouts() {
    fmt.Println("\nSelect with multiple timeouts:")
    
    ch := make(chan string)
    
    // Try to receive with different timeouts
    for i := 0; i < 3; i++ {
        select {
        case msg := <-ch:
            fmt.Printf("Received: %s\n", msg)
        case <-time.After(50 * time.Millisecond):
            fmt.Println("Short timeout")
        case <-time.After(100 * time.Millisecond):
            fmt.Println("Long timeout")
        }
    }
}

// Select with channel closure
func selectWithChannelClosure() {
    fmt.Println("\nSelect with channel closure:")
    
    ch := make(chan int)
    done := make(chan bool)
    
    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
            time.Sleep(50 * time.Millisecond)
        }
        close(ch)
        done <- true
    }()
    
    for {
        select {
        case val, ok := <-ch:
            if !ok {
                fmt.Println("Channel closed")
                return
            }
            fmt.Printf("Received: %d\n", val)
        case <-done:
            fmt.Println("Done signal received")
            return
        }
    }
}

func main() {
    basicSelect()
    selectWithDefault()
    selectWithTimeout()
    selectMultipleCases()
    selectWithSend()
    selectWithMultipleTimeouts()
    selectWithChannelClosure()
} 