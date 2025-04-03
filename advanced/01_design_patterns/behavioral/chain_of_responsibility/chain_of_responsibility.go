package chain_of_responsibility

// Handler defines the interface for handling requests
type Handler interface {
    SetNext(Handler) Handler
    Handle(string) string
}

// BaseHandler provides default implementation for Handler
type BaseHandler struct {
    next Handler
}

// SetNext sets the next handler in the chain
func (h *BaseHandler) SetNext(next Handler) Handler {
    h.next = next
    return next
}

// Handle passes the request to the next handler
func (h *BaseHandler) Handle(request string) string {
    if h.next != nil {
        return h.next.Handle(request)
    }
    return ""
}

// ConcreteHandlerA handles requests it is responsible for
type ConcreteHandlerA struct {
    *BaseHandler
}

// NewConcreteHandlerA creates a new ConcreteHandlerA
func NewConcreteHandlerA() *ConcreteHandlerA {
    return &ConcreteHandlerA{
        BaseHandler: &BaseHandler{},
    }
}

// Handle implements the Handler interface
func (h *ConcreteHandlerA) Handle(request string) string {
    if request == "A" {
        return "ConcreteHandlerA handled the request"
    }
    return h.BaseHandler.Handle(request)
}

// ConcreteHandlerB handles requests it is responsible for
type ConcreteHandlerB struct {
    *BaseHandler
}

// NewConcreteHandlerB creates a new ConcreteHandlerB
func NewConcreteHandlerB() *ConcreteHandlerB {
    return &ConcreteHandlerB{
        BaseHandler: &BaseHandler{},
    }
}

// Handle implements the Handler interface
func (h *ConcreteHandlerB) Handle(request string) string {
    if request == "B" {
        return "ConcreteHandlerB handled the request"
    }
    return h.BaseHandler.Handle(request)
}

// ConcreteHandlerC handles requests it is responsible for
type ConcreteHandlerC struct {
    *BaseHandler
}

// NewConcreteHandlerC creates a new ConcreteHandlerC
func NewConcreteHandlerC() *ConcreteHandlerC {
    return &ConcreteHandlerC{
        BaseHandler: &BaseHandler{},
    }
}

// Handle implements the Handler interface
func (h *ConcreteHandlerC) Handle(request string) string {
    if request == "C" {
        return "ConcreteHandlerC handled the request"
    }
    return h.BaseHandler.Handle(request)
}

// Client represents a client that uses the chain of responsibility
type Client struct {
    chain Handler
}

// NewClient creates a new Client
func NewClient(chain Handler) *Client {
    return &Client{chain: chain}
}

// SendRequest sends a request to the chain
func (c *Client) SendRequest(request string) string {
    return c.chain.Handle(request)
} 