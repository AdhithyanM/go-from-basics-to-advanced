package state

import "testing"

func TestConcreteStateA(t *testing.T) {
    context := NewContext(NewConcreteStateA())
    
    // Test initial state
    if _, ok := context.state.(*ConcreteStateA); !ok {
        t.Error("Expected initial state to be ConcreteStateA")
    }
    
    // Test state transition
    result := context.Request()
    expected := "State A handled the request and transitioned to State B"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Verify state transition
    if _, ok := context.state.(*ConcreteStateB); !ok {
        t.Error("Expected state to be ConcreteStateB after transition")
    }
}

func TestConcreteStateB(t *testing.T) {
    context := NewContext(NewConcreteStateB())
    
    // Test initial state
    if _, ok := context.state.(*ConcreteStateB); !ok {
        t.Error("Expected initial state to be ConcreteStateB")
    }
    
    // Test state transition
    result := context.Request()
    expected := "State B handled the request and transitioned to State A"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Verify state transition
    if _, ok := context.state.(*ConcreteStateA); !ok {
        t.Error("Expected state to be ConcreteStateA after transition")
    }
}

func TestContext(t *testing.T) {
    context := NewContext(NewConcreteStateA())
    
    // Test initial state
    if _, ok := context.state.(*ConcreteStateA); !ok {
        t.Error("Expected initial state to be ConcreteStateA")
    }
    
    // Test state setting
    context.SetState(NewConcreteStateB())
    if _, ok := context.state.(*ConcreteStateB); !ok {
        t.Error("Expected state to be ConcreteStateB after setting")
    }
    
    // Test request handling
    result := context.Request()
    expected := "State B handled the request and transitioned to State A"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    client := NewClient(NewConcreteStateA())
    
    // Test initial state
    if state := client.GetCurrentState(); state != "StateA" {
        t.Errorf("Expected initial state to be 'StateA', got '%s'", state)
    }
    
    // Test request handling
    result := client.MakeRequest()
    expected := "State A handled the request and transitioned to State B"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Verify state transition
    if state := client.GetCurrentState(); state != "StateB" {
        t.Errorf("Expected state to be 'StateB', got '%s'", state)
    }
    
    // Test another request
    result = client.MakeRequest()
    expected = "State B handled the request and transitioned to State A"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Verify state transition
    if state := client.GetCurrentState(); state != "StateA" {
        t.Errorf("Expected state to be 'StateA', got '%s'", state)
    }
} 