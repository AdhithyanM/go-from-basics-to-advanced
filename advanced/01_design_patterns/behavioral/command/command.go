package command

type Command interface {
    Execute() string
    Undo() string
}

type Receiver struct {
    state string
}

func NewReceiver() *Receiver {
    return &Receiver{state: "initial"}
}

func (r *Receiver) Action(param string) string {
    r.state = param
    return "Receiver: " + param
}

func (r *Receiver) GetState() string {
    return r.state
}

type ConcreteCommandA struct {
    receiver *Receiver
    param    string
}

func NewConcreteCommandA(receiver *Receiver, param string) *ConcreteCommandA {
    return &ConcreteCommandA{receiver: receiver, param: param}
}

func (c *ConcreteCommandA) Execute() string {
    return c.receiver.Action(c.param)
}

func (c *ConcreteCommandA) Undo() string {
    return c.receiver.Action("initial")
}

type ConcreteCommandB struct {
    receiver *Receiver
    param    string
}

func NewConcreteCommandB(receiver *Receiver, param string) *ConcreteCommandB {
    return &ConcreteCommandB{receiver: receiver, param: param}
}

func (c *ConcreteCommandB) Execute() string {
    return c.receiver.Action(c.param)
}

func (c *ConcreteCommandB) Undo() string {
    return c.receiver.Action("initial")
}

type Invoker struct {
    history []Command
}

func NewInvoker() *Invoker {
    return &Invoker{history: make([]Command, 0)}
}

func (i *Invoker) ExecuteCommand(cmd Command) string {
    result := cmd.Execute()
    i.history = append(i.history, cmd)
    return result
}

func (i *Invoker) UndoLastCommand() string {
    if len(i.history) == 0 {
        return "No commands to undo"
    }
    lastCmd := i.history[len(i.history)-1]
    i.history = i.history[:len(i.history)-1]
    return lastCmd.Undo()
}

type Client struct {
    invoker  *Invoker
    receiver *Receiver
}

func NewClient() *Client {
    return &Client{
        invoker:  NewInvoker(),
        receiver: NewReceiver(),
    }
}

func (c *Client) RunCommand(param string) string {
    cmd := NewConcreteCommandA(c.receiver, param)
    return c.invoker.ExecuteCommand(cmd)
}

func (c *Client) UndoLastCommand() string {
    return c.invoker.UndoLastCommand()
} 