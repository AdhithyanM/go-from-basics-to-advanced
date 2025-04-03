package flyweight

import (
	"sync"
)

// FlyweightFactory creates and manages flyweight objects
type FlyweightFactory struct {
    flyweights map[string]*ConcreteFlyweight
    mutex      sync.RWMutex
}

// NewFlyweightFactory creates a new FlyweightFactory
func NewFlyweightFactory() *FlyweightFactory {
    return &FlyweightFactory{
        flyweights: make(map[string]*ConcreteFlyweight),
    }
}

// GetFlyweight returns a flyweight object based on the key
func (f *FlyweightFactory) GetFlyweight(key string) *ConcreteFlyweight {
    f.mutex.RLock()
    if flyweight, exists := f.flyweights[key]; exists {
        f.mutex.RUnlock()
        return flyweight
    }
    f.mutex.RUnlock()

    f.mutex.Lock()
    defer f.mutex.Unlock()

    // Double-check after acquiring write lock
    if flyweight, exists := f.flyweights[key]; exists {
        return flyweight
    }

    flyweight := NewConcreteFlyweight(key)
    f.flyweights[key] = flyweight
    return flyweight
}

// ConcreteFlyweight represents the shared flyweight object
type ConcreteFlyweight struct {
    intrinsicState string
}

// NewConcreteFlyweight creates a new ConcreteFlyweight
func NewConcreteFlyweight(intrinsicState string) *ConcreteFlyweight {
    return &ConcreteFlyweight{
        intrinsicState: intrinsicState,
    }
}

// Operation performs an operation using both intrinsic and extrinsic state
func (f *ConcreteFlyweight) Operation(extrinsicState string) string {
    return "Intrinsic: " + f.intrinsicState + ", Extrinsic: " + extrinsicState
}

// UnsharedConcreteFlyweight represents unshared flyweight objects
type UnsharedConcreteFlyweight struct {
    allState string
}

// NewUnsharedConcreteFlyweight creates a new UnsharedConcreteFlyweight
func NewUnsharedConcreteFlyweight(allState string) *UnsharedConcreteFlyweight {
    return &UnsharedConcreteFlyweight{
        allState: allState,
    }
}

// Operation performs an operation using the unshared state
func (f *UnsharedConcreteFlyweight) Operation() string {
    return "Unshared: " + f.allState
}

// Client represents a client that uses flyweight objects
type Client struct {
    factory *FlyweightFactory
}

// NewClient creates a new Client
func NewClient(factory *FlyweightFactory) *Client {
    return &Client{
        factory: factory,
    }
}

// UseFlyweight uses a flyweight object
func (c *Client) UseFlyweight(key string, extrinsicState string) string {
    flyweight := c.factory.GetFlyweight(key)
    return flyweight.Operation(extrinsicState)
}

// UseUnsharedFlyweight uses an unshared flyweight object
func (c *Client) UseUnsharedFlyweight(allState string) string {
    flyweight := NewUnsharedConcreteFlyweight(allState)
    return flyweight.Operation()
} 