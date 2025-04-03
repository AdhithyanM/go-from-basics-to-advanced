package prototype

import "fmt"

// Prototype is the interface that declares the cloning method
type Prototype interface {
    Clone() Prototype
    GetInfo() string
}

// ConcretePrototype1 represents a concrete prototype
type ConcretePrototype1 struct {
    name string
    data map[string]string
}

// NewConcretePrototype1 creates a new ConcretePrototype1
func NewConcretePrototype1(name string) *ConcretePrototype1 {
    return &ConcretePrototype1{
        name: name,
        data: make(map[string]string),
    }
}

// Clone creates a deep copy of the prototype
func (p *ConcretePrototype1) Clone() Prototype {
    // Create a new prototype
    clone := &ConcretePrototype1{
        name: p.name,
        data: make(map[string]string),
    }
    
    // Copy the data map
    for k, v := range p.data {
        clone.data[k] = v
    }
    
    return clone
}

// GetInfo returns information about the prototype
func (p *ConcretePrototype1) GetInfo() string {
    return fmt.Sprintf("ConcretePrototype1: %s, Data: %v", p.name, p.data)
}

// SetData sets data in the prototype
func (p *ConcretePrototype1) SetData(key, value string) {
    p.data[key] = value
}

// ConcretePrototype2 represents another concrete prototype
type ConcretePrototype2 struct {
    name string
    data []int
}

// NewConcretePrototype2 creates a new ConcretePrototype2
func NewConcretePrototype2(name string) *ConcretePrototype2 {
    return &ConcretePrototype2{
        name: name,
        data: make([]int, 0),
    }
}

// Clone creates a deep copy of the prototype
func (p *ConcretePrototype2) Clone() Prototype {
    // Create a new prototype
    clone := &ConcretePrototype2{
        name: p.name,
        data: make([]int, len(p.data)),
    }
    
    // Copy the data slice
    copy(clone.data, p.data)
    
    return clone
}

// GetInfo returns information about the prototype
func (p *ConcretePrototype2) GetInfo() string {
    return fmt.Sprintf("ConcretePrototype2: %s, Data: %v", p.name, p.data)
}

// AddData adds data to the prototype
func (p *ConcretePrototype2) AddData(value int) {
    p.data = append(p.data, value)
}

// PrototypeRegistry manages a collection of prototypes
type PrototypeRegistry struct {
    prototypes map[string]Prototype
}

// NewPrototypeRegistry creates a new PrototypeRegistry
func NewPrototypeRegistry() *PrototypeRegistry {
    return &PrototypeRegistry{
        prototypes: make(map[string]Prototype),
    }
}

// AddPrototype adds a prototype to the registry
func (r *PrototypeRegistry) AddPrototype(key string, prototype Prototype) {
    r.prototypes[key] = prototype
}

// GetPrototype retrieves a prototype from the registry
func (r *PrototypeRegistry) GetPrototype(key string) Prototype {
    if prototype, exists := r.prototypes[key]; exists {
        return prototype.Clone()
    }
    return nil
}

// RemovePrototype removes a prototype from the registry
func (r *PrototypeRegistry) RemovePrototype(key string) {
    delete(r.prototypes, key)
} 