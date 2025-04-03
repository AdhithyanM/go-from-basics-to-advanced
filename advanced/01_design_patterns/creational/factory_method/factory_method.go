package factory_method

// Product interface defines the common interface for all products.
// All concrete products must implement these methods.
type Product interface {
    // Operation performs the product's main functionality
    Operation() string
    // GetName returns the product's name
    GetName() string
}

// ConcreteProductA represents a specific type of product.
// It implements the Product interface with its own behavior.
type ConcreteProductA struct {
    // name stores the product's name
    name string
}

// Operation implements the Product interface for ConcreteProductA
func (p *ConcreteProductA) Operation() string {
    return "ConcreteProductA operation"
}

// GetName returns the product's name
func (p *ConcreteProductA) GetName() string {
    return p.name
}

// ConcreteProductB represents another type of product.
// It implements the Product interface with different behavior.
type ConcreteProductB struct {
    // name stores the product's name
    name string
}

// Operation implements the Product interface for ConcreteProductB
func (p *ConcreteProductB) Operation() string {
    return "ConcreteProductB operation"
}

// GetName returns the product's name
func (p *ConcreteProductB) GetName() string {
    return p.name
}

// Creator interface defines the factory method interface.
// Concrete creators implement this interface to create specific products.
type Creator interface {
    // FactoryMethod creates and returns a Product
    FactoryMethod(name string) Product
    // SomeOperation performs an operation using the created product
    SomeOperation(name string) string
}

// ConcreteCreatorA creates ConcreteProductA instances.
// It implements the Creator interface.
type ConcreteCreatorA struct{}

// FactoryMethod creates a new ConcreteProductA instance
func (c *ConcreteCreatorA) FactoryMethod(name string) Product {
    return &ConcreteProductA{name: name}
}

// SomeOperation performs an operation using the created product
func (c *ConcreteCreatorA) SomeOperation(name string) string {
    product := c.FactoryMethod(name)
    return "CreatorA: " + product.Operation()
}

// ConcreteCreatorB creates ConcreteProductB instances.
// It implements the Creator interface.
type ConcreteCreatorB struct{}

// FactoryMethod creates a new ConcreteProductB instance
func (c *ConcreteCreatorB) FactoryMethod(name string) Product {
    return &ConcreteProductB{name: name}
}

// SomeOperation performs an operation using the created product
func (c *ConcreteCreatorB) SomeOperation(name string) string {
    product := c.FactoryMethod(name)
    return "CreatorB: " + product.Operation()
}

// Client represents a client that uses the factory method.
// It demonstrates how to use creators and products.
type Client struct{}

// UseCreator demonstrates the use of a creator
func (c *Client) UseCreator(creator Creator, name string) string {
    return creator.SomeOperation(name)
}

// NewCreator creates a new creator based on type.
// This is a simple factory function that returns the appropriate creator.
func NewCreator(creatorType string) Creator {
    switch creatorType {
    case "A":
        return &ConcreteCreatorA{}
    case "B":
        return &ConcreteCreatorB{}
    default:
        return nil
    }
} 