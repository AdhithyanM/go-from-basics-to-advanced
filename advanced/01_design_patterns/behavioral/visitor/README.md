# Visitor Pattern

The Visitor pattern represents an operation to be performed on elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.

## Key Components

- **Visitor**: Interface that declares visit operations for each element type
- **ConcreteVisitor**: Implements each operation declared by Visitor
- **Element**: Interface that declares an Accept operation
- **ConcreteElement**: Implements the Accept operation
- **ObjectStructure**: Can enumerate its elements and provide a high-level interface

## Implementation Features

- Separation of algorithm from object structure
- Easy addition of new operations
- Accumulation of state
- Breaking encapsulation when needed
- Type-safe implementation

## Use Cases

- Abstract syntax trees
- Document object models
- Complex object structures
- Report generation
- Data structure traversal
- Compiler design
- XML/JSON processing
- UI component traversal

## When to Use

- When you need to perform operations on all elements of a complex object structure
- When the object structure classes rarely change but operations change frequently
- When you want to avoid polluting classes with operation code
- When related operations need to be grouped
- When you need to maintain state during traversal

## When Not to Use

- When object structure classes change frequently
- When operations rarely change
- When visitor operations need access to private members
- When the object structure is simple
- When breaking encapsulation is not acceptable

## Best Practices

1. **Visitor Interface Design**

   - Keep visit methods focused
   - Use meaningful method names
   - Consider return types carefully
   - Document visit behavior

2. **Element Interface Design**

   - Keep Accept method simple
   - Consider making elements immutable
   - Document element responsibilities
   - Consider element lifecycle

3. **Object Structure Design**

   - Provide clear traversal methods
   - Consider iteration order
   - Handle null elements
   - Document structure constraints

4. **Implementation**
   - Keep visitors stateless if possible
   - Handle errors gracefully
   - Consider performance implications
   - Document dependencies

## Example Usage

```go
// Create elements and structure
client := NewClient()
client.AddElement(NewConcreteElementA("ElementA"))
client.AddElement(NewConcreteElementB("ElementB"))

// Create and execute visitors
visitor1 := NewConcreteVisitor1()
results := client.ExecuteVisitor(visitor1)
// Output: ["Visitor1: Processing ElementA", "Visitor1: Processing ElementB"]

visitor2 := NewConcreteVisitor2()
results = client.ExecuteVisitor(visitor2)
// Output: ["Visitor2: Processing ElementA", "Visitor2: Processing ElementB"]
```

## Testing

Run the tests using:

```bash
go test -v
```

The test suite includes:

- Testing concrete elements
- Testing concrete visitors
- Testing object structure
- Testing client usage
- Testing different visitor implementations

## Common Pitfalls

1. **Design Issues**

   - Too many visitor methods
   - Complex visitor hierarchies
   - Tight coupling to elements
   - Poor separation of concerns

2. **Implementation Issues**

   - Breaking encapsulation
   - State management problems
   - Type safety issues
   - Performance bottlenecks

3. **Maintenance Issues**

   - Hard to add new element types
   - Complex visitor implementations
   - Poor documentation
   - Unclear dependencies

4. **Usage Issues**
   - Incorrect traversal order
   - Memory leaks
   - Thread safety problems
   - Error handling complexity

## Related Patterns

- **Iterator**: Often used with Visitor to traverse structures
- **Composite**: Often provides the structure for Visitor
- **Interpreter**: Can use Visitor to interpret expressions
- **Command**: Can use Visitor to parameterize operations
- **Observer**: Can use Visitor to update multiple objects
