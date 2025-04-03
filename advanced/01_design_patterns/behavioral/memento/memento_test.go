package memento

import "testing"

func TestOriginator(t *testing.T) {
    originator := NewOriginator("initial")
    if state := originator.GetState(); state != "initial" {
        t.Errorf("Expected state 'initial', got '%s'", state)
    }

    originator.SetState("new")
    if state := originator.GetState(); state != "new" {
        t.Errorf("Expected state 'new', got '%s'", state)
    }

    memento := originator.CreateMemento()
    if state := memento.GetState(); state != "new" {
        t.Errorf("Expected memento state 'new', got '%s'", state)
    }

    originator.RestoreMemento(memento)
    if state := originator.GetState(); state != "new" {
        t.Errorf("Expected restored state 'new', got '%s'", state)
    }
}

func TestMemento(t *testing.T) {
    memento := NewMemento("test")
    if state := memento.GetState(); state != "test" {
        t.Errorf("Expected state 'test', got '%s'", state)
    }
}

func TestCaretaker(t *testing.T) {
    caretaker := NewCaretaker()
    memento1 := NewMemento("state1")
    memento2 := NewMemento("state2")

    caretaker.AddMemento(memento1)
    caretaker.AddMemento(memento2)

    if m := caretaker.GetMemento(0); m.GetState() != "state1" {
        t.Errorf("Expected state 'state1', got '%s'", m.GetState())
    }

    if m := caretaker.GetMemento(1); m.GetState() != "state2" {
        t.Errorf("Expected state 'state2', got '%s'", m.GetState())
    }

    if m := caretaker.GetMemento(2); m != nil {
        t.Error("Expected nil for out of bounds index")
    }
}

func TestClient(t *testing.T) {
    client := NewClient("initial")
    if state := client.GetCurrentState(); state != "initial" {
        t.Errorf("Expected state 'initial', got '%s'", state)
    }

    // Save initial state
    client.SaveState()

    // Change state
    client.SetState("new")
    if state := client.GetCurrentState(); state != "new" {
        t.Errorf("Expected state 'new', got '%s'", state)
    }

    // Save new state
    client.SaveState()

    // Restore to initial state
    client.RestoreState(0)
    if state := client.GetCurrentState(); state != "initial" {
        t.Errorf("Expected restored state 'initial', got '%s'", state)
    }

    // Restore to new state
    client.RestoreState(1)
    if state := client.GetCurrentState(); state != "new" {
        t.Errorf("Expected restored state 'new', got '%s'", state)
    }

    // Try to restore invalid state
    client.RestoreState(2)
    if state := client.GetCurrentState(); state != "new" {
        t.Errorf("Expected state to remain 'new', got '%s'", state)
    }
} 