# Template Method Pattern

The Template Method pattern defines the skeleton of an algorithm in a method, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

## Key Components

- **AbstractClass**: Defines the template method and abstract operations
- **ConcreteClass**: Implements the primitive operations
- **Template Method**: Defines the algorithm skeleton
- **Primitive Operations**: Abstract methods implemented by subclasses
- **Hook Methods**: Optional methods that subclasses may override

## Implementation Features

- Algorithm structure defined in base class
- Common code factored into a single location
- Customizable steps through subclassing
- Optional hook methods
- Consistent algorithm structure across implementations

## Use Cases

- Framework development
- Algorithm standardization
- Code reuse
- Library development
- Application workflows
- Document processing
- Data transformation pipelines
- Event handling systems

## When to Use

- When you want to implement invariant parts of an algorithm once
- When you want to localize common behavior in a single class
- When you want to control subclass extensions
- When you have multiple classes with similar algorithms
- When you want to avoid code duplication

## When Not to Use

- When algorithm steps vary significantly
- When algorithm structure is not fixed
- When subclassing would be too complex
- When flexibility in algorithm structure is needed
- When code reuse through inheritance is not appropriate

## Best Practices

1. **Template Method Design**

   - Keep the template method focused
   - Minimize the number of primitive operations
   - Document the algorithm structure
   - Use meaningful method names

2. **Hook Methods**

   - Keep hooks optional
   - Provide default implementations
   - Document hook purposes
   - Use hooks sparingly

3. **Primitive Operations**

   - Keep operations simple
   - Make operations cohesive
   - Document required behavior
   - Consider operation granularity

4. **Inheritance Structure**
   - Keep inheritance hierarchy shallow
   - Document extension points
   - Consider composition alternatives
   - Use interfaces when appropriate

## Example Usage

```go
// Create concrete classes
classA := NewConcreteClassA()
classB := NewConcreteClassB()

// Execute template method
resultA := classA.TemplateMethod()
// Output: "ConcreteClassA: Step 1 -> ConcreteClassA: Step 2 -> ConcreteClassA: Hook"

resultB := classB.TemplateMethod()
// Output: "ConcreteClassB: Step 1 -> ConcreteClassB: Step 2"
```

## Testing

Run the tests using:

```bash
go test -v
```

The test suite includes:

- Testing primitive operations
- Testing hook methods
- Testing template method
- Testing different concrete classes
- Testing variations in algorithm execution

## Common Pitfalls

1. **Inheritance Issues**

   - Deep inheritance hierarchies
   - Tight coupling to base class
   - Inflexible algorithm structure
   - Poor extension points

2. **Method Design**

   - Too many primitive operations
   - Complex template methods
   - Unclear hook purposes
   - Poor method naming

3. **Code Organization**

   - Duplicate code in subclasses
   - Inconsistent method implementations
   - Poor documentation
   - Unclear algorithm structure

4. **Maintenance**
   - Hard to modify algorithm structure
   - Difficult to add new variations
   - Complex testing requirements
   - Poor version compatibility

## Related Patterns

- **Strategy**: Encapsulates interchangeable algorithms
- **Factory Method**: Often used in template methods
- **Builder**: Can use template method to construct objects
- **Abstract Factory**: Can use template method to create families of objects
- **Command**: Can use template method to structure command execution
