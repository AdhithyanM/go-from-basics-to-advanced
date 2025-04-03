package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Basic atomic counter
func atomicCounter() {
    fmt.Println("Basic atomic counter:")
    
    var counter int64
    
    // Launch multiple goroutines that increment counter
    for i := 0; i < 5; i++ {
        go func() {
            for j := 0; j < 3; j++ {
                atomic.AddInt64(&counter, 1)
                fmt.Printf("Counter: %d\n", atomic.LoadInt64(&counter))
                time.Sleep(50 * time.Millisecond)
            }
        }()
    }
    
    time.Sleep(200 * time.Millisecond)
    fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
}

// Atomic boolean flag
func atomicBoolean() {
    fmt.Println("\nAtomic boolean flag:")
    
    var flag uint32
    
    // Set flag
    atomic.StoreUint32(&flag, 1)
    fmt.Printf("Flag value: %v\n", atomic.LoadUint32(&flag) == 1)
    
    // Compare and swap
    swapped := atomic.CompareAndSwapUint32(&flag, 1, 0)
    fmt.Printf("Swapped: %v, New flag value: %v\n", swapped, atomic.LoadUint32(&flag) == 1)
}

// Atomic pointer
func atomicPointer() {
    fmt.Println("\nAtomic pointer:")
    
    type Data struct {
        value int
    }
    
    var ptr atomic.Value
    
    // Store pointer
    data := &Data{value: 42}
    ptr.Store(data)
    
    // Load pointer
    if loaded, ok := ptr.Load().(*Data); ok {
        fmt.Printf("Loaded value: %d\n", loaded.value)
    }
}

// Atomic operations with different types
func atomicTypes() {
    fmt.Println("\nAtomic operations with different types:")
    
    // Int32
    var i32 int32
    atomic.AddInt32(&i32, 1)
    fmt.Printf("Int32: %d\n", atomic.LoadInt32(&i32))
    
    // Uint32
    var u32 uint32
    atomic.AddUint32(&u32, 1)
    fmt.Printf("Uint32: %d\n", atomic.LoadUint32(&u32))
    
    // Int64
    var i64 int64
    atomic.AddInt64(&i64, 1)
    fmt.Printf("Int64: %d\n", atomic.LoadInt64(&i64))
    
    // Uint64
    var u64 uint64
    atomic.AddUint64(&u64, 1)
    fmt.Printf("Uint64: %d\n", atomic.LoadUint64(&u64))
}

// Atomic value with custom type
func atomicCustomType() {
    fmt.Println("\nAtomic value with custom type:")
    
    type Config struct {
        Host     string
        Port     int
        Enabled  bool
    }
    
    var config atomic.Value
    
    // Store initial config
    initialConfig := Config{
        Host:    "localhost",
        Port:    8080,
        Enabled: true,
    }
    config.Store(initialConfig)
    
    // Load and modify config
    if currentConfig, ok := config.Load().(Config); ok {
        fmt.Printf("Current config: %+v\n", currentConfig)
        
        // Update config
        newConfig := Config{
            Host:    "example.com",
            Port:    9090,
            Enabled: false,
        }
        config.Store(newConfig)
        
        // Load updated config
        if updatedConfig, ok := config.Load().(Config); ok {
            fmt.Printf("Updated config: %+v\n", updatedConfig)
        }
    }
}

// Atomic operations with CAS
func atomicCAS() {
    fmt.Println("\nAtomic operations with CAS:")
    
    var value int64 = 42
    
    // Try to update value using CAS
    for {
        current := atomic.LoadInt64(&value)
        newValue := current + 1
        
        if atomic.CompareAndSwapInt64(&value, current, newValue) {
            fmt.Printf("Successfully updated value to: %d\n", newValue)
            break
        } else {
            fmt.Println("Failed to update value, retrying...")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() {
    atomicCounter()
    atomicBoolean()
    atomicPointer()
    atomicTypes()
    atomicCustomType()
    atomicCAS()
} 