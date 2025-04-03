package command

import "testing"

func TestReceiver(t *testing.T) {
    receiver := NewReceiver()
    if state := receiver.GetState(); state != "initial" {
        t.Errorf("Expected state 'initial', got '%s'", state)
    }

    expected := "Receiver: test"
    if result := receiver.Action("test"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    if state := receiver.GetState(); state != "test" {
        t.Errorf("Expected state 'test', got '%s'", state)
    }
}

func TestConcreteCommandA(t *testing.T) {
    receiver := NewReceiver()
    cmd := NewConcreteCommandA(receiver, "test")

    expected := "Receiver: test"
    if result := cmd.Execute(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    expected = "Receiver: initial"
    if result := cmd.Undo(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteCommandB(t *testing.T) {
    receiver := NewReceiver()
    cmd := NewConcreteCommandB(receiver, "test")

    expected := "Receiver: test"
    if result := cmd.Execute(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    expected = "Receiver: initial"
    if result := cmd.Undo(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestInvoker(t *testing.T) {
    invoker := NewInvoker()
    receiver := NewReceiver()

    cmdA := NewConcreteCommandA(receiver, "A")
    cmdB := NewConcreteCommandB(receiver, "B")

    // Execute first command
    expected := "Receiver: A"
    if result := invoker.ExecuteCommand(cmdA); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    // Execute second command
    expected = "Receiver: B"
    if result := invoker.ExecuteCommand(cmdB); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    // Undo last command
    expected = "Receiver: initial"
    if result := invoker.UndoLastCommand(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    // Undo first command
    expected = "Receiver: initial"
    if result := invoker.UndoLastCommand(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    // Try to undo when no commands left
    expected = "No commands to undo"
    if result := invoker.UndoLastCommand(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    client := NewClient()

    // Run command
    expected := "Receiver: test"
    if result := client.RunCommand("test"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }

    // Undo command
    expected = "Receiver: initial"
    if result := client.UndoLastCommand(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 