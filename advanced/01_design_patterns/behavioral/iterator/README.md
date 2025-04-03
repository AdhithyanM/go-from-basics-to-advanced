# Iterator Pattern

The Iterator pattern provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

## Implementation Details

### Key Components

1. **Iterator**

   - Defines the interface for accessing and traversing elements
   - Declares methods for checking if there are more elements
   - Declares methods for accessing the next element

2. **Aggregate**

   - Defines the interface for creating iterators
   - Declares a method for creating an iterator

3. **ConcreteIterator**

   - Implements the Iterator interface
   - Keeps track of the current position in the traversal
   - Implements the traversal algorithm

4. **ConcreteAggregate**
   - Implements the Aggregate interface
   - Returns an instance of the ConcreteIterator
   - Stores the collection of elements

### Implementation Features

- **Encapsulation**: Hides the internal structure of the collection
- **Flexibility**: Different traversal algorithms can be implemented
- **Single Responsibility**: Separates traversal logic from collection
- **Polymorphism**: Works with any collection that implements the interface

## Use Cases

1. **Collection Traversal**

   - Iterating over arrays, lists, or trees
   - Accessing elements in a specific order
   - Filtering or transforming elements

2. **Data Structures**

   - Custom data structures
   - Complex collections
   - Hierarchical structures

3. **Stream Processing**
   - Processing data streams
   - Lazy evaluation
   - Infinite sequences

## When to Use

- When you need to access elements of a collection without exposing its structure
- When you need multiple traversal algorithms for the same collection
- When you want to provide a uniform interface for traversing different collections
- When you want to decouple the collection from its traversal

## When Not to Use

- When the collection is simple and direct access is sufficient
- When performance is critical and the overhead is unacceptable
- When the collection structure is unlikely to change
- When the traversal algorithm is fixed

## Best Practices

1. **Implementation**

   - Keep iterators simple and focused
   - Make iterators immutable
   - Consider using the Strategy pattern for different traversal algorithms
   - Handle concurrent modifications

2. **Testing**

   - Test each iterator independently
   - Test different traversal scenarios
   - Test edge cases
   - Test concurrent access

3. **Documentation**
   - Document the traversal order
   - Explain any limitations
   - Provide usage examples
   - Document thread safety

## Example Usage

```go
// Create client
client := NewClient()

// Add items
client.AddItems(1, 2, 3)

// Iterate over items
result := client.Iterate()
```

## Testing

Run the tests with:

```bash
go test -v
```

Check test coverage:

```bash
go test -cover
```

Run benchmarks:

```bash
go test -bench=.
```

## Common Pitfalls

1. **Complexity**

   - Overly complex iterators
   - Complex state management
   - Nested iterations

2. **Performance**

   - Iterator object overhead
   - Memory usage
   - Traversal efficiency

3. **Maintenance**
   - Iterator state management
   - Concurrent modifications
   - Error handling

## Related Patterns

- **Composite**: Treats individual objects and compositions uniformly
- **Visitor**: Separates algorithms from object structures
- **Strategy**: Encapsulates algorithms in separate classes
