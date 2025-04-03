# Testing in Go

This directory contains examples of different testing approaches in Go, demonstrating various testing methodologies and best practices.

## Directory Structure

- `basic/`: Basic unit testing examples
- `table_driven/`: Table-driven test examples
- `benchmarks/`: Benchmark testing examples
- `coverage/`: Test coverage examples
- `parallel/`: Parallel testing examples
- `testify/`: Examples using the testify package

## Basic Testing (`basic/`)

Demonstrates fundamental unit testing concepts:

- Basic test structure
- Testing function outputs
- Error handling tests
- Testing struct methods
- Testing state changes

Example:

```go
func TestAdd(t *testing.T) {
    calc := &Calculator{}
    result := calc.Add(2, 3)
    expected := 5.0

    if result != expected {
        t.Errorf("Add(2, 3) = %f; expected %f", result, expected)
    }
}
```

## Table-Driven Tests (`table_driven/`)

Shows how to write table-driven tests, which are more maintainable and comprehensive:

- Multiple test cases in a single test
- Structured test data
- Consistent test patterns
- Easy to add new test cases

Example:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Benchmarks (`benchmarks/`)

Demonstrates performance testing:

- Basic benchmarking
- Benchmarking with setup
- Sub-benchmarks
- Memory allocation reporting
- Parallel benchmarking
- Different input size testing

Example:

```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fibonacci(20)
    }
}
```

## Coverage (`coverage/`)

Shows how to achieve and measure test coverage:

- Comprehensive test cases
- Edge case handling
- Error scenarios
- State management
- Coverage reporting

Run with coverage:

```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Parallel Testing (`parallel/`)

Demonstrates parallel test execution:

- Parallel test execution
- Shared resource handling
- Cleanup functions
- Subtest parallelization
- Race condition detection

Example:

```go
func TestParallel(t *testing.T) {
    t.Run("subtest", func(t *testing.T) {
        t.Parallel()
        // Test implementation
    })
}
```

Run with race detection:

```bash
go test -race ./...
```

## Testify Examples (`testify/`)

Shows how to use the testify package for:

- Assertions
- Mocking
- Suite testing
- HTTP testing
- Test helpers

## Running Tests

Basic test execution:

```bash
go test ./...
```

Verbose output:

```bash
go test -v ./...
```

Run specific test:

```bash
go test -run TestName ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Best Practices

1. **Test Organization**

   - One test file per source file
   - Test files end with `_test.go`
   - Tests in the same package as the code being tested

2. **Test Naming**

   - Test functions start with `Test`
   - Test names should be descriptive
   - Use subtests for related test cases

3. **Test Coverage**

   - Aim for high test coverage
   - Focus on critical paths
   - Test edge cases and error conditions

4. **Table-Driven Tests**

   - Use for multiple test cases
   - Keep test data separate from test logic
   - Use descriptive test case names

5. **Benchmarking**

   - Use realistic input sizes
   - Reset timer after setup
   - Report memory allocations when relevant

6. **Parallel Testing**

   - Use `t.Parallel()` for independent tests
   - Handle shared resources carefully
   - Use race detector for concurrent code

7. **Error Handling**
   - Test both success and error cases
   - Verify error messages
   - Test edge cases and invalid inputs
