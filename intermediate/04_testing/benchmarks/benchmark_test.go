package benchmarks

import (
	"fmt"
	"testing"
)

// Example function to benchmark
func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// Example function with optimization
func fibonacciOptimized(n int) int {
    if n < 2 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

// Benchmark the recursive implementation
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fibonacci(20)
    }
}

// Benchmark the optimized implementation
func BenchmarkFibonacciOptimized(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fibonacciOptimized(20)
    }
}

// Example of a benchmark with setup
func BenchmarkWithSetup(b *testing.B) {
    // Setup code
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }

    // Reset timer to exclude setup time
    b.ResetTimer()

    // Benchmark code
    for i := 0; i < b.N; i++ {
        for j := range data {
            data[j] *= 2
        }
    }
}

// Example of a benchmark with sub-benchmarks
func BenchmarkOperations(b *testing.B) {
    b.Run("addition", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = 1 + 1
        }
    })

    b.Run("multiplication", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = 2 * 2
        }
    })
}

// Example of a benchmark with different input sizes
func BenchmarkDifferentSizes(b *testing.B) {
    sizes := []int{10, 100, 1000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
            data := make([]int, size)
            for i := range data {
                data[i] = i
            }
            
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                for j := range data {
                    data[j] *= 2
                }
            }
        })
    }
}

// Example of a benchmark with memory allocation
func BenchmarkMemoryAllocation(b *testing.B) {
    b.ReportAllocs() // Report memory allocations
    
    for i := 0; i < b.N; i++ {
        s := make([]int, 1000)
        _ = s
    }
}

// Example of a benchmark with parallel execution
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            fibonacciOptimized(20)
            b.SetBytes(1)
        }
    })
} 