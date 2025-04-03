package memento

// Originator represents the object whose state needs to be saved
type Originator struct {
    state string
}

// NewOriginator creates a new Originator
func NewOriginator(state string) *Originator {
    return &Originator{state: state}
}

// SetState sets the state of the originator
func (o *Originator) SetState(state string) {
    o.state = state
}

// GetState returns the current state
func (o *Originator) GetState() string {
    return o.state
}

// CreateMemento creates a memento containing the current state
func (o *Originator) CreateMemento() *Memento {
    return NewMemento(o.state)
}

// RestoreMemento restores the state from a memento
func (o *Originator) RestoreMemento(m *Memento) {
    o.state = m.GetState()
}

// Memento stores the state of the Originator
type Memento struct {
    state string
}

// NewMemento creates a new Memento
func NewMemento(state string) *Memento {
    return &Memento{state: state}
}

// GetState returns the stored state
func (m *Memento) GetState() string {
    return m.state
}

// Caretaker is responsible for the memento's safekeeping
type Caretaker struct {
    mementos []*Memento
}

// NewCaretaker creates a new Caretaker
func NewCaretaker() *Caretaker {
    return &Caretaker{
        mementos: make([]*Memento, 0),
    }
}

// AddMemento adds a memento to the caretaker
func (c *Caretaker) AddMemento(m *Memento) {
    c.mementos = append(c.mementos, m)
}

// GetMemento retrieves a memento by index
func (c *Caretaker) GetMemento(index int) *Memento {
    if index < 0 || index >= len(c.mementos) {
        return nil
    }
    return c.mementos[index]
}

// Client represents a client that uses the memento pattern
type Client struct {
    originator *Originator
    caretaker  *Caretaker
}

// NewClient creates a new Client
func NewClient(initialState string) *Client {
    return &Client{
        originator: NewOriginator(initialState),
        caretaker:  NewCaretaker(),
    }
}

// SaveState saves the current state
func (c *Client) SaveState() {
    memento := c.originator.CreateMemento()
    c.caretaker.AddMemento(memento)
}

// RestoreState restores a previous state
func (c *Client) RestoreState(index int) {
    memento := c.caretaker.GetMemento(index)
    if memento != nil {
        c.originator.RestoreMemento(memento)
    }
}

// GetCurrentState returns the current state
func (c *Client) GetCurrentState() string {
    return c.originator.GetState()
}

// SetState sets the current state
func (c *Client) SetState(state string) {
    c.originator.SetState(state)
} 