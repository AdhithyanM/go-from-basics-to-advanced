package adapter

// Target defines the domain-specific interface that Client uses
type Target interface {
    Request() string
}

// Adaptee contains some useful behavior, but its interface is incompatible
// with the existing client code
type Adaptee struct {
    specificRequest string
}

// NewAdaptee creates a new Adaptee with the given specific request
func NewAdaptee(specificRequest string) *Adaptee {
    return &Adaptee{
        specificRequest: specificRequest,
    }
}

// SpecificRequest returns the specific request of the Adaptee
func (a *Adaptee) SpecificRequest() string {
    return a.specificRequest
}

// Adapter makes the Adaptee's interface compatible with the Target's interface
type Adapter struct {
    adaptee *Adaptee
}

// NewAdapter creates a new Adapter that wraps the Adaptee
func NewAdapter(adaptee *Adaptee) *Adapter {
    return &Adapter{
        adaptee: adaptee,
    }
}

// Request implements the Target interface by adapting the Adaptee's SpecificRequest
func (a *Adapter) Request() string {
    return "Adapter: " + a.adaptee.SpecificRequest()
}

// Client represents a client that uses the Target interface
type Client struct{}

// UseTarget uses the Target interface without knowing the concrete implementation
func (c *Client) UseTarget(target Target) string {
    return target.Request()
}

// ClassAdapter adapts the Adaptee through inheritance
type ClassAdapter struct {
    *Adaptee
}

// NewClassAdapter creates a new ClassAdapter
func NewClassAdapter(specificRequest string) *ClassAdapter {
    return &ClassAdapter{
        Adaptee: NewAdaptee(specificRequest),
    }
}

// Request implements the Target interface
func (a *ClassAdapter) Request() string {
    return a.SpecificRequest()
}

// ObjectAdapter adapts the Adaptee through composition
type ObjectAdapter struct {
    adaptee *Adaptee
}

// NewObjectAdapter creates a new ObjectAdapter
func NewObjectAdapter(specificRequest string) *ObjectAdapter {
    return &ObjectAdapter{
        adaptee: NewAdaptee(specificRequest),
    }
}

// Request implements the Target interface
func (a *ObjectAdapter) Request() string {
    return a.adaptee.SpecificRequest()
}

// Example usage of different adapter types
func Example() {
    client := &Client{}
    
    // Using class adapter
    classAdapter := NewClassAdapter("Class Adapter Data")
    classResult := client.UseTarget(classAdapter)
    
    // Using object adapter
    objectAdapter := NewObjectAdapter("Object Adapter Data")
    objectResult := client.UseTarget(objectAdapter)
    
    _ = classResult   // Use the result
    _ = objectResult // Use the result
} 