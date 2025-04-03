package decorator

// Component defines the interface for objects that can have responsibilities added to them
type Component interface {
    Operation() string
}

// ConcreteComponent defines an object to which additional responsibilities can be attached
type ConcreteComponent struct {
    name string
}

// NewConcreteComponent creates a new ConcreteComponent
func NewConcreteComponent(name string) *ConcreteComponent {
    return &ConcreteComponent{name: name}
}

// Operation implements the Component interface
func (c *ConcreteComponent) Operation() string {
    return "ConcreteComponent " + c.name
}

// Decorator maintains a reference to a Component object and defines an interface that conforms to Component's interface
type Decorator struct {
    component Component
}

// NewDecorator creates a new Decorator
func NewDecorator(component Component) *Decorator {
    return &Decorator{component: component}
}

// Operation implements the Component interface
func (d *Decorator) Operation() string {
    return d.component.Operation()
}

// ConcreteDecoratorA adds responsibilities to the component
type ConcreteDecoratorA struct {
    *Decorator
    addedState string
}

// NewConcreteDecoratorA creates a new ConcreteDecoratorA
func NewConcreteDecoratorA(component Component, addedState string) *ConcreteDecoratorA {
    return &ConcreteDecoratorA{
        Decorator:   NewDecorator(component),
        addedState: addedState,
    }
}

// Operation implements the Component interface
func (d *ConcreteDecoratorA) Operation() string {
    return d.Decorator.Operation() + " with ConcreteDecoratorA " + d.addedState
}

// ConcreteDecoratorB adds responsibilities to the component
type ConcreteDecoratorB struct {
    *Decorator
    addedBehavior func(string) string
}

// NewConcreteDecoratorB creates a new ConcreteDecoratorB
func NewConcreteDecoratorB(component Component, addedBehavior func(string) string) *ConcreteDecoratorB {
    return &ConcreteDecoratorB{
        Decorator:     NewDecorator(component),
        addedBehavior: addedBehavior,
    }
}

// Operation implements the Component interface
func (d *ConcreteDecoratorB) Operation() string {
    return d.addedBehavior(d.Decorator.Operation())
} 