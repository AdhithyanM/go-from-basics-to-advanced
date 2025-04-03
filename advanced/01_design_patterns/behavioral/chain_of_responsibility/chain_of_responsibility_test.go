package chain_of_responsibility

import "testing"

func TestConcreteHandlerA(t *testing.T) {
    handler := NewConcreteHandlerA()
    expected := "ConcreteHandlerA handled the request"
    if result := handler.Handle("A"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    if result := handler.Handle("B"); result != "" {
        t.Errorf("Expected empty string, got '%s'", result)
    }
}

func TestConcreteHandlerB(t *testing.T) {
    handler := NewConcreteHandlerB()
    expected := "ConcreteHandlerB handled the request"
    if result := handler.Handle("B"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    if result := handler.Handle("A"); result != "" {
        t.Errorf("Expected empty string, got '%s'", result)
    }
}

func TestConcreteHandlerC(t *testing.T) {
    handler := NewConcreteHandlerC()
    expected := "ConcreteHandlerC handled the request"
    if result := handler.Handle("C"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    if result := handler.Handle("A"); result != "" {
        t.Errorf("Expected empty string, got '%s'", result)
    }
}

func TestChain(t *testing.T) {
    // Create handlers
    handlerA := NewConcreteHandlerA()
    handlerB := NewConcreteHandlerB()
    handlerC := NewConcreteHandlerC()

    // Build the chain
    handlerA.SetNext(handlerB).SetNext(handlerC)

    // Test each handler in the chain
    tests := []struct {
        request  string
        expected string
    }{
        {"A", "ConcreteHandlerA handled the request"},
        {"B", "ConcreteHandlerB handled the request"},
        {"C", "ConcreteHandlerC handled the request"},
        {"D", ""},
    }

    for _, test := range tests {
        if result := handlerA.Handle(test.request); result != test.expected {
            t.Errorf("For request '%s', expected '%s', got '%s'", test.request, test.expected, result)
        }
    }
}

func TestClient(t *testing.T) {
    // Create handlers
    handlerA := NewConcreteHandlerA()
    handlerB := NewConcreteHandlerB()
    handlerC := NewConcreteHandlerC()

    // Build the chain
    handlerA.SetNext(handlerB).SetNext(handlerC)

    // Create client
    client := NewClient(handlerA)

    // Test client
    expected := "ConcreteHandlerB handled the request"
    if result := client.SendRequest("B"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 