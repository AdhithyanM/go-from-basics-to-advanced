package visitor

// Element defines the interface for elements that can be visited
type Element interface {
    Accept(visitor Visitor) string
}

// Visitor defines the interface for visitors
type Visitor interface {
    VisitConcreteElementA(element *ConcreteElementA) string
    VisitConcreteElementB(element *ConcreteElementB) string
}

// ConcreteElementA implements the Element interface
type ConcreteElementA struct {
    name string
}

// NewConcreteElementA creates a new ConcreteElementA
func NewConcreteElementA(name string) *ConcreteElementA {
    return &ConcreteElementA{
        name: name,
    }
}

// Accept implements the Element interface
func (e *ConcreteElementA) Accept(visitor Visitor) string {
    return visitor.VisitConcreteElementA(e)
}

// GetName returns the element's name
func (e *ConcreteElementA) GetName() string {
    return e.name
}

// ConcreteElementB implements the Element interface
type ConcreteElementB struct {
    name string
}

// NewConcreteElementB creates a new ConcreteElementB
func NewConcreteElementB(name string) *ConcreteElementB {
    return &ConcreteElementB{
        name: name,
    }
}

// Accept implements the Element interface
func (e *ConcreteElementB) Accept(visitor Visitor) string {
    return visitor.VisitConcreteElementB(e)
}

// GetName returns the element's name
func (e *ConcreteElementB) GetName() string {
    return e.name
}

// ConcreteVisitor1 implements the Visitor interface
type ConcreteVisitor1 struct{}

// NewConcreteVisitor1 creates a new ConcreteVisitor1
func NewConcreteVisitor1() *ConcreteVisitor1 {
    return &ConcreteVisitor1{}
}

// VisitConcreteElementA implements visitor behavior for ConcreteElementA
func (v *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) string {
    return "Visitor1: Processing " + element.GetName()
}

// VisitConcreteElementB implements visitor behavior for ConcreteElementB
func (v *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) string {
    return "Visitor1: Processing " + element.GetName()
}

// ConcreteVisitor2 implements the Visitor interface
type ConcreteVisitor2 struct{}

// NewConcreteVisitor2 creates a new ConcreteVisitor2
func NewConcreteVisitor2() *ConcreteVisitor2 {
    return &ConcreteVisitor2{}
}

// VisitConcreteElementA implements visitor behavior for ConcreteElementA
func (v *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) string {
    return "Visitor2: Processing " + element.GetName()
}

// VisitConcreteElementB implements visitor behavior for ConcreteElementB
func (v *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) string {
    return "Visitor2: Processing " + element.GetName()
}

// ObjectStructure maintains a collection of elements
type ObjectStructure struct {
    elements []Element
}

// NewObjectStructure creates a new ObjectStructure
func NewObjectStructure() *ObjectStructure {
    return &ObjectStructure{
        elements: make([]Element, 0),
    }
}

// AddElement adds an element to the structure
func (o *ObjectStructure) AddElement(element Element) {
    o.elements = append(o.elements, element)
}

// Accept visits all elements with the given visitor
func (o *ObjectStructure) Accept(visitor Visitor) []string {
    results := make([]string, len(o.elements))
    for i, element := range o.elements {
        results[i] = element.Accept(visitor)
    }
    return results
}

// Client represents a client that uses the visitor pattern
type Client struct {
    structure *ObjectStructure
}

// NewClient creates a new Client
func NewClient() *Client {
    return &Client{
        structure: NewObjectStructure(),
    }
}

// AddElement adds an element to the structure
func (c *Client) AddElement(element Element) {
    c.structure.AddElement(element)
}

// ExecuteVisitor executes a visitor on all elements
func (c *Client) ExecuteVisitor(visitor Visitor) []string {
    return c.structure.Accept(visitor)
} 