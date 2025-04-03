package main

import (
	"fmt"
	"time"
)

// Basic channel operations
func basicChannels() {
    fmt.Println("Basic channel operations:")
    
    // Create an unbuffered channel
    ch := make(chan string)
    
    // Send in goroutine
    go func() {
        ch <- "Hello from channel!"
    }()
    
    // Receive
    msg := <-ch
    fmt.Println(msg)
}

// Buffered channels
func bufferedChannels() {
    fmt.Println("\nBuffered channels:")
    
    // Create a buffered channel with capacity 2
    ch := make(chan int, 2)
    
    // Send values without blocking
    ch <- 1
    ch <- 2
    
    // Receive values
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

// Channel direction
func channelDirection() {
    fmt.Println("\nChannel direction:")
    
    // Create channels
    ch := make(chan int)
    
    // Send-only channel
    go func(ch chan<- int) {
        ch <- 42
    }(ch)
    
    // Receive-only channel
    go func(ch <-chan int) {
        fmt.Println(<-ch)
    }(ch)
    
    time.Sleep(50 * time.Millisecond)
}

// Channel patterns
func channelPatterns() {
    fmt.Println("\nChannel patterns:")
    
    // Generator pattern
    numbers := make(chan int)
    go func() {
        for i := 0; i < 3; i++ {
            numbers <- i
        }
        close(numbers)
    }()
    
    // Receive from closed channel
    for num := range numbers {
        fmt.Printf("Received: %d\n", num)
    }
}

// Channel with timeout
func channelTimeout() {
    fmt.Println("\nChannel with timeout:")
    
    ch := make(chan string)
    
    // Send with timeout
    select {
    case ch <- "message":
        fmt.Println("Message sent")
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout sending message")
    }
    
    // Receive with timeout
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout receiving message")
    }
}

// Channel with multiple goroutines
func multipleGoroutines() {
    fmt.Println("\nMultiple goroutines with channels:")
    
    ch := make(chan int)
    
    // Producer
    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(ch)
    }()
    
    // Multiple consumers
    for i := 0; i < 2; i++ {
        go func(id int) {
            for num := range ch {
                fmt.Printf("Consumer %d received: %d\n", id, num)
            }
        }(i)
    }
    
    time.Sleep(400 * time.Millisecond)
}

// Channel with select
func channelSelect() {
    fmt.Println("\nChannel with select:")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Send on both channels
    go func() {
        ch1 <- "First channel"
    }()
    
    go func() {
        ch2 <- "Second channel"
    }()
    
    // Select from multiple channels
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Printf("Received from ch1: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("Received from ch2: %s\n", msg2)
        }
    }
}

func main() {
    basicChannels()
    bufferedChannels()
    channelDirection()
    channelPatterns()
    channelTimeout()
    multipleGoroutines()
    channelSelect()
} 