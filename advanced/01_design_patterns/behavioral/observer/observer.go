package observer

// Observer defines the interface for objects that should be notified of changes
type Observer interface {
    Update(subject Subject)
}

// Subject defines the interface for objects that are observed
type Subject interface {
    RegisterObserver(observer Observer)
    RemoveObserver(observer Observer)
    NotifyObservers()
}

// ConcreteSubject implements the Subject interface
type ConcreteSubject struct {
    observers []Observer
    state     string
}

// NewConcreteSubject creates a new ConcreteSubject
func NewConcreteSubject() *ConcreteSubject {
    return &ConcreteSubject{
        observers: make([]Observer, 0),
    }
}

// RegisterObserver adds an observer to the list
func (s *ConcreteSubject) RegisterObserver(observer Observer) {
    s.observers = append(s.observers, observer)
}

// RemoveObserver removes an observer from the list
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
    for i, obs := range s.observers {
        if obs == observer {
            s.observers = append(s.observers[:i], s.observers[i+1:]...)
            break
        }
    }
}

// NotifyObservers notifies all observers of a change
func (s *ConcreteSubject) NotifyObservers() {
    for _, observer := range s.observers {
        observer.Update(s)
    }
}

// SetState sets the state and notifies observers
func (s *ConcreteSubject) SetState(state string) {
    s.state = state
    s.NotifyObservers()
}

// GetState returns the current state
func (s *ConcreteSubject) GetState() string {
    return s.state
}

// ConcreteObserverA implements the Observer interface
type ConcreteObserverA struct {
    state string
}

// NewConcreteObserverA creates a new ConcreteObserverA
func NewConcreteObserverA() *ConcreteObserverA {
    return &ConcreteObserverA{}
}

// Update updates the observer's state
func (o *ConcreteObserverA) Update(subject Subject) {
    if s, ok := subject.(*ConcreteSubject); ok {
        o.state = s.GetState()
    }
}

// GetState returns the current state
func (o *ConcreteObserverA) GetState() string {
    return o.state
}

// ConcreteObserverB implements the Observer interface
type ConcreteObserverB struct {
    state string
}

// NewConcreteObserverB creates a new ConcreteObserverB
func NewConcreteObserverB() *ConcreteObserverB {
    return &ConcreteObserverB{}
}

// Update updates the observer's state
func (o *ConcreteObserverB) Update(subject Subject) {
    if s, ok := subject.(*ConcreteSubject); ok {
        o.state = s.GetState()
    }
}

// GetState returns the current state
func (o *ConcreteObserverB) GetState() string {
    return o.state
}

// Client represents a client that uses the observer pattern
type Client struct {
    subject   *ConcreteSubject
    observers []Observer
}

// NewClient creates a new Client
func NewClient() *Client {
    return &Client{
        subject:   NewConcreteSubject(),
        observers: make([]Observer, 0),
    }
}

// AddObserver adds an observer to the subject
func (c *Client) AddObserver(observer Observer) {
    c.subject.RegisterObserver(observer)
    c.observers = append(c.observers, observer)
}

// RemoveObserver removes an observer from the subject
func (c *Client) RemoveObserver(observer Observer) {
    c.subject.RemoveObserver(observer)
    for i, obs := range c.observers {
        if obs == observer {
            c.observers = append(c.observers[:i], c.observers[i+1:]...)
            break
        }
    }
}

// SetState sets the subject's state
func (c *Client) SetState(state string) {
    c.subject.SetState(state)
}

// GetObserverState returns the state of an observer
func (c *Client) GetObserverState(index int) string {
    if index < 0 || index >= len(c.observers) {
        return ""
    }
    if obs, ok := c.observers[index].(*ConcreteObserverA); ok {
        return obs.GetState()
    }
    if obs, ok := c.observers[index].(*ConcreteObserverB); ok {
        return obs.GetState()
    }
    return ""
} 