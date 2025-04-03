# Decorator Pattern

The Decorator pattern attaches additional responsibilities to an object dynamically. It provides a flexible alternative to subclassing for extending functionality.

## Implementation Details

### Key Components

1. **Component**

   - Defines the interface for objects that can have responsibilities added to them
   - Declares the operation that can be decorated

2. **ConcreteComponent**

   - Defines an object to which additional responsibilities can be attached
   - Implements the Component interface

3. **Decorator**

   - Maintains a reference to a Component object
   - Defines an interface that conforms to Component's interface
   - Forwards requests to the Component object

4. **ConcreteDecorator**
   - Adds responsibilities to the component
   - Can add state or behavior
   - Extends the Decorator class

### Implementation Features

- **Dynamic**: Responsibilities can be added at runtime
- **Flexible**: Multiple decorators can be combined
- **Transparent**: Decorators are transparent to the client
- **Extensible**: New decorators can be added without changing existing code

## Use Cases

1. **GUI Components**

   - Adding borders to windows
   - Adding scrollbars to text areas
   - Adding tooltips to buttons

2. **I/O Streams**

   - Adding buffering to file streams
   - Adding compression to network streams
   - Adding encryption to data streams

3. **Middleware**
   - Adding logging to service calls
   - Adding caching to database queries
   - Adding authentication to API endpoints

## When to Use

- When you need to add responsibilities to individual objects dynamically
- When you want to add responsibilities that can be withdrawn
- When extension by subclassing is impractical
- When you need to add responsibilities to objects without affecting other objects

## When Not to Use

- When the number of decorators would be excessive
- When the decorators would make the code too complex
- When the performance overhead is unacceptable
- When the responsibilities are not truly independent

## Best Practices

1. **Implementation**

   - Keep the Component interface simple
   - Make decorators transparent to the client
   - Use composition over inheritance
   - Consider using the Strategy pattern for complex behaviors

2. **Testing**

   - Test individual decorators
   - Test combinations of decorators
   - Test the order of decoration
   - Verify proper forwarding of operations

3. **Documentation**
   - Document the decoration process
   - Explain the responsibilities of each decorator
   - Provide examples of common combinations
   - Document any limitations or constraints

## Example Usage

```go
// Create a component
component := NewConcreteComponent("A")

// Add decorators
decoratorA := NewConcreteDecoratorA(component, "state1")
decoratorB := NewConcreteDecoratorB(decoratorA, func(s string) string {
    return "{" + s + "}"
})

// Use the decorated component
result := decoratorB.Operation()
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

   - Too many decorators
   - Complex decoration chains
   - Unclear responsibility boundaries

2. **Performance**

   - Multiple layers of indirection
   - Memory overhead from decorators
   - Costly operations in decorators

3. **Maintenance**
   - Difficulty in understanding decoration chains
   - Hard to debug decorated objects
   - Complex state management

## Related Patterns

- **Composite**: Decorator is often used with Composite
- **Strategy**: Can be used to implement complex behaviors in decorators
- **Proxy**: Similar structure but different purpose
