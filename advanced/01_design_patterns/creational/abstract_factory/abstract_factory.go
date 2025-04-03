package abstract_factory

// ProductA is the interface for the first product family
type ProductA interface {
    OperationA() string
}

// ProductB is the interface for the second product family
type ProductB interface {
    OperationB() string
    InteractWithA(a ProductA) string
}

// AbstractFactory is the interface for the abstract factory
type AbstractFactory interface {
    CreateProductA() ProductA
    CreateProductB() ProductB
}

// ConcreteProductA1 is a concrete implementation of ProductA
type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) OperationA() string {
    return "ConcreteProductA1 operation"
}

// ConcreteProductB1 is a concrete implementation of ProductB
type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) OperationB() string {
    return "ConcreteProductB1 operation"
}

func (p *ConcreteProductB1) InteractWithA(a ProductA) string {
    return "ConcreteProductB1 interacting with " + a.OperationA()
}

// ConcreteProductA2 is another concrete implementation of ProductA
type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) OperationA() string {
    return "ConcreteProductA2 operation"
}

// ConcreteProductB2 is another concrete implementation of ProductB
type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) OperationB() string {
    return "ConcreteProductB2 operation"
}

func (p *ConcreteProductB2) InteractWithA(a ProductA) string {
    return "ConcreteProductB2 interacting with " + a.OperationA()
}

// ConcreteFactory1 creates products of the first family
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA {
    return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
    return &ConcreteProductB1{}
}

// ConcreteFactory2 creates products of the second family
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() ProductA {
    return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
    return &ConcreteProductB2{}
}

// Client represents a client that uses the abstract factory
type Client struct{}

// CreateProducts creates and uses products from a factory
func (c *Client) CreateProducts(factory AbstractFactory) (string, string) {
    productA := factory.CreateProductA()
    productB := factory.CreateProductB()
    return productA.OperationA(), productB.InteractWithA(productA)
}

// NewFactory creates a new factory based on type
func NewFactory(factoryType string) AbstractFactory {
    switch factoryType {
    case "1":
        return &ConcreteFactory1{}
    case "2":
        return &ConcreteFactory2{}
    default:
        return nil
    }
} 