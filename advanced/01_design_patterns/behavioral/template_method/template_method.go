package template_method

// AbstractClass defines the template method and primitive operations
type AbstractClass interface {
    TemplateMethod() string
    PrimitiveOperation1() string
    PrimitiveOperation2() string
    Hook() string
}

// AbstractClassImpl provides default implementation for the template method
type AbstractClassImpl struct {
    concrete AbstractClass
}

// NewAbstractClassImpl creates a new AbstractClassImpl
func NewAbstractClassImpl(concrete AbstractClass) *AbstractClassImpl {
    return &AbstractClassImpl{
        concrete: concrete,
    }
}

// TemplateMethod defines the algorithm skeleton
func (a *AbstractClassImpl) TemplateMethod() string {
    result := a.concrete.PrimitiveOperation1()
    result += " -> " + a.concrete.PrimitiveOperation2()
    if hook := a.concrete.Hook(); hook != "" {
        result += " -> " + hook
    }
    return result
}

// ConcreteClassA implements the AbstractClass interface
type ConcreteClassA struct {
    *AbstractClassImpl
}

// NewConcreteClassA creates a new ConcreteClassA
func NewConcreteClassA() *ConcreteClassA {
    concrete := &ConcreteClassA{}
    concrete.AbstractClassImpl = NewAbstractClassImpl(concrete)
    return concrete
}

// PrimitiveOperation1 implements a step of the algorithm
func (c *ConcreteClassA) PrimitiveOperation1() string {
    return "ConcreteClassA: Step 1"
}

// PrimitiveOperation2 implements a step of the algorithm
func (c *ConcreteClassA) PrimitiveOperation2() string {
    return "ConcreteClassA: Step 2"
}

// Hook provides additional behavior
func (c *ConcreteClassA) Hook() string {
    return "ConcreteClassA: Hook"
}

// ConcreteClassB implements the AbstractClass interface
type ConcreteClassB struct {
    *AbstractClassImpl
}

// NewConcreteClassB creates a new ConcreteClassB
func NewConcreteClassB() *ConcreteClassB {
    concrete := &ConcreteClassB{}
    concrete.AbstractClassImpl = NewAbstractClassImpl(concrete)
    return concrete
}

// PrimitiveOperation1 implements a step of the algorithm
func (c *ConcreteClassB) PrimitiveOperation1() string {
    return "ConcreteClassB: Step 1"
}

// PrimitiveOperation2 implements a step of the algorithm
func (c *ConcreteClassB) PrimitiveOperation2() string {
    return "ConcreteClassB: Step 2"
}

// Hook provides additional behavior (empty implementation)
func (c *ConcreteClassB) Hook() string {
    return ""
}

// Client represents a client that uses the template method pattern
type Client struct {
    class AbstractClass
}

// NewClient creates a new Client
func NewClient(class AbstractClass) *Client {
    return &Client{
        class: class,
    }
}

// ExecuteAlgorithm executes the template method
func (c *Client) ExecuteAlgorithm() string {
    return c.class.TemplateMethod()
} 