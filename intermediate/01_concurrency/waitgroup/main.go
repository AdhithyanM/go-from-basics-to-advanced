package main

import (
	"fmt"
	"sync"
	"time"
)

// Basic WaitGroup usage
func basicWaitGroup() {
    fmt.Println("Basic WaitGroup usage:")
    
    var wg sync.WaitGroup
    
    // Launch multiple goroutines
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d finished\n", num)
        }(i)
    }
    
    // Wait for all goroutines to complete
    wg.Wait()
    fmt.Println("All goroutines completed")
}

// WaitGroup with error handling
func waitGroupWithErrors() {
    fmt.Println("\nWaitGroup with error handling:")
    
    var wg sync.WaitGroup
    errChan := make(chan error, 3)
    
    // Launch goroutines that might fail
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            
            if num == 1 {
                errChan <- fmt.Errorf("error in goroutine %d", num)
                return
            }
            
            fmt.Printf("Goroutine %d completed successfully\n", num)
        }(i)
    }
    
    // Wait for all goroutines to complete
    wg.Wait()
    close(errChan)
    
    // Check for errors
    for err := range errChan {
        fmt.Printf("Error: %v\n", err)
    }
}

// WaitGroup with timeout
func waitGroupWithTimeout() {
    fmt.Println("\nWaitGroup with timeout:")
    
    var wg sync.WaitGroup
    done := make(chan struct{})
    
    // Launch long-running goroutines
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            time.Sleep(100 * time.Millisecond)
            fmt.Printf("Goroutine %d finished\n", num)
        }(i)
    }
    
    // Wait with timeout
    go func() {
        wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        fmt.Println("All goroutines completed")
    case <-time.After(50 * time.Millisecond):
        fmt.Println("Timeout waiting for goroutines")
    }
}

// WaitGroup with multiple stages
func waitGroupWithStages() {
    fmt.Println("\nWaitGroup with multiple stages:")
    
    var wg1, wg2 sync.WaitGroup
    
    // First stage
    for i := 0; i < 2; i++ {
        wg1.Add(1)
        go func(num int) {
            defer wg1.Done()
            fmt.Printf("Stage 1 - Goroutine %d\n", num)
        }(i)
    }
    
    // Wait for first stage
    wg1.Wait()
    
    // Second stage
    for i := 0; i < 2; i++ {
        wg2.Add(1)
        go func(num int) {
            defer wg2.Done()
            fmt.Printf("Stage 2 - Goroutine %d\n", num)
        }(i)
    }
    
    // Wait for second stage
    wg2.Wait()
    fmt.Println("All stages completed")
}

// WaitGroup with shared resource
func waitGroupWithSharedResource() {
    fmt.Println("\nWaitGroup with shared resource:")
    
    var wg sync.WaitGroup
    var counter int
    var mu sync.Mutex
    
    // Launch multiple goroutines that modify shared resource
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            
            mu.Lock()
            counter++
            fmt.Printf("Goroutine %d: counter = %d\n", num, counter)
            mu.Unlock()
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter)
}

// WaitGroup with cleanup
func waitGroupWithCleanup() {
    fmt.Println("\nWaitGroup with cleanup:")
    
    var wg sync.WaitGroup
    resources := make([]string, 3)
    
    // Launch goroutines that use resources
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            defer func() {
                // Cleanup
                resources[num] = ""
                fmt.Printf("Cleaned up resource %d\n", num)
            }()
            
            resources[num] = fmt.Sprintf("Resource %d", num)
            fmt.Printf("Using resource %d\n", num)
            time.Sleep(50 * time.Millisecond)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All resources cleaned up")
}

func main() {
    basicWaitGroup()
    waitGroupWithErrors()
    waitGroupWithTimeout()
    waitGroupWithStages()
    waitGroupWithSharedResource()
    waitGroupWithCleanup()
} 