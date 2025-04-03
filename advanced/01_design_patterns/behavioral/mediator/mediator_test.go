package mediator

import "testing"

func TestConcreteColleagueA(t *testing.T) {
    colleague := NewConcreteColleagueA("A")
    mediator := NewConcreteMediator()
    colleague.SetMediator(mediator)

    expected := "A received: test"
    if result := colleague.Receive("test"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteColleagueB(t *testing.T) {
    colleague := NewConcreteColleagueB("B")
    mediator := NewConcreteMediator()
    colleague.SetMediator(mediator)

    expected := "B received: test"
    if result := colleague.Receive("test"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestMediator(t *testing.T) {
    mediator := NewConcreteMediator()
    colleagueA := NewConcreteColleagueA("A")
    colleagueB := NewConcreteColleagueB("B")

    mediator.RegisterColleague("A", colleagueA)
    mediator.RegisterColleague("B", colleagueB)

    // Test notification
    colleagueA.Send("test")
    // The message should be received by colleagueB
    expected := "B received: test"
    if result := colleagueB.Receive("test"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    client := NewClient()
    colleagueA := NewConcreteColleagueA("A")
    colleagueB := NewConcreteColleagueB("B")

    client.AddColleague("A", colleagueA)
    client.AddColleague("B", colleagueB)

    // Test message sending
    client.SendMessage(colleagueA, "test")
    // The message should be received by colleagueB
    expected := "B received: test"
    if result := colleagueB.Receive("test"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 