package state

// State defines the interface for state-specific behavior
type State interface {
    Handle(context *Context) string
}

// Context maintains a reference to the current state
type Context struct {
    state State
}

// NewContext creates a new Context with an initial state
func NewContext(initialState State) *Context {
    return &Context{
        state: initialState,
    }
}

// SetState changes the current state
func (c *Context) SetState(state State) {
    c.state = state
}

// Request delegates the request to the current state
func (c *Context) Request() string {
    return c.state.Handle(c)
}

// ConcreteStateA implements the State interface
type ConcreteStateA struct{}

// NewConcreteStateA creates a new ConcreteStateA
func NewConcreteStateA() *ConcreteStateA {
    return &ConcreteStateA{}
}

// Handle implements state-specific behavior
func (s *ConcreteStateA) Handle(context *Context) string {
    context.SetState(NewConcreteStateB())
    return "State A handled the request and transitioned to State B"
}

// ConcreteStateB implements the State interface
type ConcreteStateB struct{}

// NewConcreteStateB creates a new ConcreteStateB
func NewConcreteStateB() *ConcreteStateB {
    return &ConcreteStateB{}
}

// Handle implements state-specific behavior
func (s *ConcreteStateB) Handle(context *Context) string {
    context.SetState(NewConcreteStateA())
    return "State B handled the request and transitioned to State A"
}

// Client represents a client that uses the state pattern
type Client struct {
    context *Context
}

// NewClient creates a new Client with an initial state
func NewClient(initialState State) *Client {
    return &Client{
        context: NewContext(initialState),
    }
}

// MakeRequest makes a request to the context
func (c *Client) MakeRequest() string {
    return c.context.Request()
}

// GetCurrentState returns the current state type
func (c *Client) GetCurrentState() string {
    switch c.context.state.(type) {
    case *ConcreteStateA:
        return "StateA"
    case *ConcreteStateB:
        return "StateB"
    default:
        return "Unknown"
    }
} 