package strategy

import "testing"

func TestConcreteStrategyA(t *testing.T) {
    strategy := NewConcreteStrategyA()
    
    result := strategy.Execute("test")
    expected := "Strategy A: test"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteStrategyB(t *testing.T) {
    strategy := NewConcreteStrategyB()
    
    result := strategy.Execute("test")
    expected := "Strategy B: test"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestContext(t *testing.T) {
    strategyA := NewConcreteStrategyA()
    context := NewContext(strategyA)
    
    // Test initial strategy
    result := context.ExecuteStrategy("test")
    expected := "Strategy A: test"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test strategy change
    strategyB := NewConcreteStrategyB()
    context.SetStrategy(strategyB)
    
    result = context.ExecuteStrategy("test")
    expected = "Strategy B: test"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    strategyA := NewConcreteStrategyA()
    client := NewClient(strategyA)
    
    // Test initial strategy
    result := client.ExecuteStrategy("test")
    expected := "Strategy A: test"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test strategy change
    strategyB := NewConcreteStrategyB()
    client.ChangeStrategy(strategyB)
    
    result = client.ExecuteStrategy("test")
    expected = "Strategy B: test"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

// TestDynamicStrategyChange tests changing strategies multiple times
func TestDynamicStrategyChange(t *testing.T) {
    client := NewClient(NewConcreteStrategyA())
    
    // Test sequence of strategy changes
    testCases := []struct {
        strategy Strategy
        input    string
        expected string
    }{
        {NewConcreteStrategyA(), "test1", "Strategy A: test1"},
        {NewConcreteStrategyB(), "test2", "Strategy B: test2"},
        {NewConcreteStrategyA(), "test3", "Strategy A: test3"},
        {NewConcreteStrategyB(), "test4", "Strategy B: test4"},
    }
    
    for _, tc := range testCases {
        client.ChangeStrategy(tc.strategy)
        result := client.ExecuteStrategy(tc.input)
        if result != tc.expected {
            t.Errorf("Expected '%s', got '%s'", tc.expected, result)
        }
    }
} 