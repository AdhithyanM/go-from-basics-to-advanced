package strategy

// Strategy defines the interface for all supported algorithms
type Strategy interface {
    Execute(data string) string
}

// Context maintains a reference to a Strategy object
type Context struct {
    strategy Strategy
}

// NewContext creates a new Context with a strategy
func NewContext(strategy Strategy) *Context {
    return &Context{
        strategy: strategy,
    }
}

// SetStrategy allows changing the strategy at runtime
func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}

// ExecuteStrategy delegates the work to the strategy
func (c *Context) ExecuteStrategy(data string) string {
    return c.strategy.Execute(data)
}

// ConcreteStrategyA implements the Strategy interface
type ConcreteStrategyA struct{}

// NewConcreteStrategyA creates a new ConcreteStrategyA
func NewConcreteStrategyA() *ConcreteStrategyA {
    return &ConcreteStrategyA{}
}

// Execute implements the algorithm using Strategy A
func (s *ConcreteStrategyA) Execute(data string) string {
    return "Strategy A: " + data
}

// ConcreteStrategyB implements the Strategy interface
type ConcreteStrategyB struct{}

// NewConcreteStrategyB creates a new ConcreteStrategyB
func NewConcreteStrategyB() *ConcreteStrategyB {
    return &ConcreteStrategyB{}
}

// Execute implements the algorithm using Strategy B
func (s *ConcreteStrategyB) Execute(data string) string {
    return "Strategy B: " + data
}

// Client represents a client that uses the strategy pattern
type Client struct {
    context *Context
}

// NewClient creates a new Client with a strategy
func NewClient(strategy Strategy) *Client {
    return &Client{
        context: NewContext(strategy),
    }
}

// ExecuteStrategy executes the current strategy
func (c *Client) ExecuteStrategy(data string) string {
    return c.context.ExecuteStrategy(data)
}

// ChangeStrategy changes the current strategy
func (c *Client) ChangeStrategy(strategy Strategy) {
    c.context.SetStrategy(strategy)
} 