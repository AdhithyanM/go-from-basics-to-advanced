package main

import (
	"fmt"
	"sync"
	"time"
)

// Basic mutex usage
func basicMutex() {
    fmt.Println("Basic mutex usage:")
    
    var mu sync.Mutex
    var counter int
    
    // Launch multiple goroutines that modify shared resource
    for i := 0; i < 5; i++ {
        go func(num int) {
            mu.Lock()
            counter++
            fmt.Printf("Goroutine %d: counter = %d\n", num, counter)
            mu.Unlock()
        }(i)
    }
    
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("Final counter value: %d\n", counter)
}

// RWMutex usage
func rwMutexExample() {
    fmt.Println("\nRWMutex usage:")
    
    var rwmu sync.RWMutex
    var data map[string]int = make(map[string]int)
    
    // Writer goroutine
    go func() {
        for i := 0; i < 3; i++ {
            rwmu.Lock()
            data[fmt.Sprintf("key%d", i)] = i
            fmt.Printf("Writer: Added key%d = %d\n", i, i)
            rwmu.Unlock()
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    // Multiple reader goroutines
    for i := 0; i < 3; i++ {
        go func(id int) {
            for j := 0; j < 3; j++ {
                rwmu.RLock()
                fmt.Printf("Reader %d: data = %v\n", id, data)
                rwmu.RUnlock()
                time.Sleep(30 * time.Millisecond)
            }
        }(i)
    }
    
    time.Sleep(200 * time.Millisecond)
}

// Mutex with defer
func mutexWithDefer() {
    fmt.Println("\nMutex with defer:")
    
    var mu sync.Mutex
    var resources []string
    
    // Function that uses mutex with defer
    addResource := func(resource string) {
        mu.Lock()
        defer mu.Unlock()
        
        resources = append(resources, resource)
        fmt.Printf("Added resource: %s\n", resource)
    }
    
    // Launch multiple goroutines
    for i := 0; i < 3; i++ {
        go func(num int) {
            addResource(fmt.Sprintf("Resource %d", num))
        }(i)
    }
    
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("Final resources: %v\n", resources)
}

// Mutex with timeout
func mutexWithTimeout() {
    fmt.Println("\nMutex with timeout:")
    
    var mu sync.Mutex
    var data int
    
    // Try to acquire lock with timeout
    go func() {
        mu.Lock()
        defer mu.Unlock()
        
        data = 42
        fmt.Println("Locked and modified data")
        time.Sleep(100 * time.Millisecond)
    }()
    
    // Try to acquire lock with timeout
    done := make(chan bool)
    go func() {
        select {
        case <-time.After(50 * time.Millisecond):
            fmt.Println("Timeout trying to acquire lock")
        default:
            mu.Lock()
            defer mu.Unlock()
            fmt.Printf("Acquired lock, data = %d\n", data)
            done <- true
        }
    }()
    
    select {
    case <-done:
        fmt.Println("Successfully acquired lock")
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout waiting for lock")
    }
}

// Mutex with deadlock prevention
func mutexWithDeadlockPrevention() {
    fmt.Println("\nMutex with deadlock prevention:")
    
    var mu1, mu2 sync.Mutex
    
    // Function that acquires locks in consistent order
    acquireLocks := func(id int) {
        if id%2 == 0 {
            mu1.Lock()
            defer mu1.Unlock()
            time.Sleep(50 * time.Millisecond)
            mu2.Lock()
            defer mu2.Unlock()
        } else {
            mu2.Lock()
            defer mu2.Unlock()
            time.Sleep(50 * time.Millisecond)
            mu1.Lock()
            defer mu1.Unlock()
        }
        fmt.Printf("Goroutine %d acquired both locks\n", id)
    }
    
    // Launch goroutines
    for i := 0; i < 3; i++ {
        go acquireLocks(i)
    }
    
    time.Sleep(200 * time.Millisecond)
}

// Mutex with try lock
func mutexWithTryLock() {
    fmt.Println("\nMutex with try lock:")
    
    var mu sync.Mutex
    
    // Try to acquire lock
    if mu.TryLock() {
        defer mu.Unlock()
        fmt.Println("Successfully acquired lock")
        time.Sleep(50 * time.Millisecond)
    } else {
        fmt.Println("Failed to acquire lock")
    }
    
    // Try again
    if mu.TryLock() {
        defer mu.Unlock()
        fmt.Println("Successfully acquired lock on second try")
    } else {
        fmt.Println("Failed to acquire lock on second try")
    }
}

func main() {
    basicMutex()
    rwMutexExample()
    mutexWithDefer()
    mutexWithTimeout()
    mutexWithDeadlockPrevention()
    mutexWithTryLock()
} 