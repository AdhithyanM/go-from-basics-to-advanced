package parallel

import (
	"sync"
	"testing"
	"time"
)

// Example function that simulates some work
func doWork(id int) {
    time.Sleep(100 * time.Millisecond)
}

// Example of parallel test execution
func TestParallel(t *testing.T) {
    tests := []struct {
        name string
        id   int
    }{
        {"test1", 1},
        {"test2", 2},
        {"test3", 3},
        {"test4", 4},
        {"test5", 5},
    }

    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Mark test as parallel
            doWork(tt.id)
        })
    }
}

// Example of parallel test with shared resources
func TestParallelWithSharedResource(t *testing.T) {
    var counter int
    var mu sync.Mutex

    tests := []struct {
        name string
        want int
    }{
        {"increment1", 1},
        {"increment2", 2},
        {"increment3", 3},
    }

    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Mark test as parallel
            
            mu.Lock()
            counter++
            got := counter
            mu.Unlock()
            
            if got != tt.want {
                t.Errorf("counter = %d; want %d", got, tt.want)
            }
        })
    }
}

// Example of parallel test with cleanup
func TestParallelWithCleanup(t *testing.T) {
    var cleanupCalled bool
    var mu sync.Mutex

    t.Cleanup(func() {
        mu.Lock()
        cleanupCalled = true
        mu.Unlock()
    })

    tests := []struct {
        name string
    }{
        {"test1"},
        {"test2"},
        {"test3"},
    }

    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Mark test as parallel
            time.Sleep(100 * time.Millisecond)
        })
    }

    // Wait for all parallel tests to complete
    time.Sleep(200 * time.Millisecond)

    mu.Lock()
    if !cleanupCalled {
        t.Error("Cleanup was not called")
    }
    mu.Unlock()
}

// Example of parallel test with subtests
func TestParallelSubtests(t *testing.T) {
    t.Run("group", func(t *testing.T) {
        tests := []struct {
            name string
            want int
        }{
            {"subtest1", 1},
            {"subtest2", 2},
            {"subtest3", 3},
        }

        for _, tt := range tests {
            tt := tt // Capture range variable
            t.Run(tt.name, func(t *testing.T) {
                t.Parallel() // Mark test as parallel
                time.Sleep(100 * time.Millisecond)
                if tt.want < 0 {
                    t.Fatal("unexpected negative value")
                }
            })
        }
    })
}

// Example of parallel test with race condition detection
func TestParallelRaceCondition(t *testing.T) {
    var counter int
    var wg sync.WaitGroup

    // Create multiple goroutines that increment the counter
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++
        }()
    }

    wg.Wait()

    if counter != 100 {
        t.Errorf("counter = %d; want 100", counter)
    }
} 