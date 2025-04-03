package mediator

// Mediator defines the interface for communication between colleagues
type Mediator interface {
    Notify(sender Colleague, event string)
}

// Colleague defines the interface for objects that communicate through the mediator
type Colleague interface {
    SetMediator(mediator Mediator)
    Send(event string)
    Receive(event string) string
}

// ConcreteMediator implements the Mediator interface
type ConcreteMediator struct {
    colleagues map[string]Colleague
}

// NewConcreteMediator creates a new ConcreteMediator
func NewConcreteMediator() *ConcreteMediator {
    return &ConcreteMediator{
        colleagues: make(map[string]Colleague),
    }
}

// RegisterColleague registers a colleague with the mediator
func (m *ConcreteMediator) RegisterColleague(name string, colleague Colleague) {
    colleague.SetMediator(m)
    m.colleagues[name] = colleague
}

// Notify handles communication between colleagues
func (m *ConcreteMediator) Notify(sender Colleague, event string) {
    for _, colleague := range m.colleagues {
        if colleague != sender {
            colleague.Receive(event)
        }
    }
}

// ConcreteColleagueA implements the Colleague interface
type ConcreteColleagueA struct {
    mediator Mediator
    name     string
}

// NewConcreteColleagueA creates a new ConcreteColleagueA
func NewConcreteColleagueA(name string) *ConcreteColleagueA {
    return &ConcreteColleagueA{name: name}
}

// SetMediator sets the mediator for the colleague
func (c *ConcreteColleagueA) SetMediator(mediator Mediator) {
    c.mediator = mediator
}

// Send sends a message through the mediator
func (c *ConcreteColleagueA) Send(event string) {
    c.mediator.Notify(c, event)
}

// Receive handles incoming messages
func (c *ConcreteColleagueA) Receive(event string) string {
    return c.name + " received: " + event
}

// ConcreteColleagueB implements the Colleague interface
type ConcreteColleagueB struct {
    mediator Mediator
    name     string
}

// NewConcreteColleagueB creates a new ConcreteColleagueB
func NewConcreteColleagueB(name string) *ConcreteColleagueB {
    return &ConcreteColleagueB{name: name}
}

// SetMediator sets the mediator for the colleague
func (c *ConcreteColleagueB) SetMediator(mediator Mediator) {
    c.mediator = mediator
}

// Send sends a message through the mediator
func (c *ConcreteColleagueB) Send(event string) {
    c.mediator.Notify(c, event)
}

// Receive handles incoming messages
func (c *ConcreteColleagueB) Receive(event string) string {
    return c.name + " received: " + event
}

// Client represents a client that uses the mediator pattern
type Client struct {
    mediator *ConcreteMediator
}

// NewClient creates a new Client
func NewClient() *Client {
    return &Client{
        mediator: NewConcreteMediator(),
    }
}

// AddColleague adds a colleague to the mediator
func (c *Client) AddColleague(name string, colleague Colleague) {
    c.mediator.RegisterColleague(name, colleague)
}

// SendMessage sends a message from one colleague to others
func (c *Client) SendMessage(sender Colleague, message string) {
    sender.Send(message)
} 