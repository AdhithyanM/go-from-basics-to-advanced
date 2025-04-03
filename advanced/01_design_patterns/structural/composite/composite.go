package composite

// Component defines the interface for objects in the composition
type Component interface {
    Operation() string
    Add(Component)
    Remove(Component)
    GetChild(int) Component
}

// Leaf represents leaf objects in the composition
type Leaf struct {
    name string
}

// NewLeaf creates a new Leaf
func NewLeaf(name string) *Leaf {
    return &Leaf{name: name}
}

// Operation implements the Component interface
func (l *Leaf) Operation() string {
    return "Leaf " + l.name
}

// Add is a no-op for Leaf
func (l *Leaf) Add(Component) {}

// Remove is a no-op for Leaf
func (l *Leaf) Remove(Component) {}

// GetChild returns nil for Leaf
func (l *Leaf) GetChild(int) Component {
    return nil
}

// Composite represents composite objects in the composition
type Composite struct {
    name     string
    children []Component
}

// NewComposite creates a new Composite
func NewComposite(name string) *Composite {
    return &Composite{
        name:     name,
        children: make([]Component, 0),
    }
}

// Operation implements the Component interface
func (c *Composite) Operation() string {
    result := "Composite " + c.name + " ["
    for i, child := range c.children {
        if i > 0 {
            result += ", "
        }
        result += child.Operation()
    }
    result += "]"
    return result
}

// Add adds a child component
func (c *Composite) Add(component Component) {
    c.children = append(c.children, component)
}

// Remove removes a child component
func (c *Composite) Remove(component Component) {
    for i, child := range c.children {
        if child == component {
            c.children = append(c.children[:i], c.children[i+1:]...)
            break
        }
    }
}

// GetChild returns a child component by index
func (c *Composite) GetChild(index int) Component {
    if index < 0 || index >= len(c.children) {
        return nil
    }
    return c.children[index]
} 