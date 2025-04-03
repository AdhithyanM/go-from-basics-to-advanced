package iterator

// Iterator defines the interface for traversing elements
type Iterator interface {
    Next() interface{}
    HasNext() bool
}

// Aggregate defines the interface for creating iterators
type Aggregate interface {
    CreateIterator() Iterator
}

// ConcreteIterator implements the Iterator interface
type ConcreteIterator struct {
    collection []interface{}
    position   int
}

// NewConcreteIterator creates a new ConcreteIterator
func NewConcreteIterator(collection []interface{}) *ConcreteIterator {
    return &ConcreteIterator{
        collection: collection,
        position:   0,
    }
}

// Next returns the next element in the collection
func (i *ConcreteIterator) Next() interface{} {
    if i.HasNext() {
        item := i.collection[i.position]
        i.position++
        return item
    }
    return nil
}

// HasNext checks if there are more elements to iterate
func (i *ConcreteIterator) HasNext() bool {
    return i.position < len(i.collection)
}

// ConcreteAggregate implements the Aggregate interface
type ConcreteAggregate struct {
    items []interface{}
}

// NewConcreteAggregate creates a new ConcreteAggregate
func NewConcreteAggregate() *ConcreteAggregate {
    return &ConcreteAggregate{
        items: make([]interface{}, 0),
    }
}

// AddItem adds an item to the aggregate
func (a *ConcreteAggregate) AddItem(item interface{}) {
    a.items = append(a.items, item)
}

// CreateIterator creates a new iterator for the aggregate
func (a *ConcreteAggregate) CreateIterator() Iterator {
    return NewConcreteIterator(a.items)
}

// Client represents a client that uses the iterator pattern
type Client struct {
    aggregate *ConcreteAggregate
}

// NewClient creates a new Client
func NewClient() *Client {
    return &Client{
        aggregate: NewConcreteAggregate(),
    }
}

// AddItems adds multiple items to the aggregate
func (c *Client) AddItems(items ...interface{}) {
    for _, item := range items {
        c.aggregate.AddItem(item)
    }
}

// Iterate iterates over the items in the aggregate
func (c *Client) Iterate() []interface{} {
    iterator := c.aggregate.CreateIterator()
    result := make([]interface{}, 0)
    for iterator.HasNext() {
        result = append(result, iterator.Next())
    }
    return result
} 