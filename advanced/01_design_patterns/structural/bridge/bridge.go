package bridge

// Implementor defines the interface for implementation classes
type Implementor interface {
    OperationImpl() string
}

// ConcreteImplementorA implements the Implementor interface
type ConcreteImplementorA struct{}

// OperationImpl implements the operation for ConcreteImplementorA
func (i *ConcreteImplementorA) OperationImpl() string {
    return "ConcreteImplementorA operation"
}

// ConcreteImplementorB implements the Implementor interface
type ConcreteImplementorB struct{}

// OperationImpl implements the operation for ConcreteImplementorB
func (i *ConcreteImplementorB) OperationImpl() string {
    return "ConcreteImplementorB operation"
}

// Abstraction defines the abstraction's interface
type Abstraction interface {
    Operation() string
}

// RefinedAbstraction extends the abstraction interface
type RefinedAbstraction struct {
    implementor Implementor
}

// NewRefinedAbstraction creates a new RefinedAbstraction
func NewRefinedAbstraction(implementor Implementor) *RefinedAbstraction {
    return &RefinedAbstraction{
        implementor: implementor,
    }
}

// Operation implements the abstraction's operation
func (a *RefinedAbstraction) Operation() string {
    return "RefinedAbstraction: " + a.implementor.OperationImpl()
}

// ExtendedAbstraction extends the abstraction further
type ExtendedAbstraction struct {
    implementor Implementor
}

// NewExtendedAbstraction creates a new ExtendedAbstraction
func NewExtendedAbstraction(implementor Implementor) *ExtendedAbstraction {
    return &ExtendedAbstraction{
        implementor: implementor,
    }
}

// Operation implements the extended abstraction's operation
func (a *ExtendedAbstraction) Operation() string {
    return "ExtendedAbstraction: " + a.implementor.OperationImpl()
}

// Client represents a client that uses the abstraction
type Client struct{}

// UseAbstraction uses the abstraction without knowing its implementation
func (c *Client) UseAbstraction(abstraction Abstraction) string {
    return abstraction.Operation()
} 