package facade

// SubsystemA represents a complex subsystem
type SubsystemA struct {
    name string
}

// NewSubsystemA creates a new SubsystemA
func NewSubsystemA(name string) *SubsystemA {
    return &SubsystemA{name: name}
}

// OperationA performs a complex operation
func (s *SubsystemA) OperationA() string {
    return "SubsystemA " + s.name + " operation"
}

// SubsystemB represents another complex subsystem
type SubsystemB struct {
    name string
}

// NewSubsystemB creates a new SubsystemB
func NewSubsystemB(name string) *SubsystemB {
    return &SubsystemB{name: name}
}

// OperationB performs a complex operation
func (s *SubsystemB) OperationB() string {
    return "SubsystemB " + s.name + " operation"
}

// SubsystemC represents yet another complex subsystem
type SubsystemC struct {
    name string
}

// NewSubsystemC creates a new SubsystemC
func NewSubsystemC(name string) *SubsystemC {
    return &SubsystemC{name: name}
}

// OperationC performs a complex operation
func (s *SubsystemC) OperationC() string {
    return "SubsystemC " + s.name + " operation"
}

// Facade provides a simplified interface to the complex subsystems
type Facade struct {
    subsystemA *SubsystemA
    subsystemB *SubsystemB
    subsystemC *SubsystemC
}

// NewFacade creates a new Facade
func NewFacade(name string) *Facade {
    return &Facade{
        subsystemA: NewSubsystemA(name + "_A"),
        subsystemB: NewSubsystemB(name + "_B"),
        subsystemC: NewSubsystemC(name + "_C"),
    }
}

// Operation provides a simplified interface to the complex subsystems
func (f *Facade) Operation() string {
    result := "Facade operation:\n"
    result += "- " + f.subsystemA.OperationA() + "\n"
    result += "- " + f.subsystemB.OperationB() + "\n"
    result += "- " + f.subsystemC.OperationC()
    return result
}

// Client represents a client that uses the Facade
type Client struct {
    facade *Facade
}

// NewClient creates a new Client
func NewClient(facade *Facade) *Client {
    return &Client{facade: facade}
}

// UseFacade uses the Facade to perform operations
func (c *Client) UseFacade() string {
    return c.facade.Operation()
} 