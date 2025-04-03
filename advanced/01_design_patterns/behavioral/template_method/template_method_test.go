package template_method

import "testing"

func TestConcreteClassA(t *testing.T) {
    classA := NewConcreteClassA()
    
    // Test primitive operations
    result := classA.PrimitiveOperation1()
    expected := "ConcreteClassA: Step 1"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    result = classA.PrimitiveOperation2()
    expected = "ConcreteClassA: Step 2"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test hook
    result = classA.Hook()
    expected = "ConcreteClassA: Hook"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test template method
    result = classA.TemplateMethod()
    expected = "ConcreteClassA: Step 1 -> ConcreteClassA: Step 2 -> ConcreteClassA: Hook"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteClassB(t *testing.T) {
    classB := NewConcreteClassB()
    
    // Test primitive operations
    result := classB.PrimitiveOperation1()
    expected := "ConcreteClassB: Step 1"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    result = classB.PrimitiveOperation2()
    expected = "ConcreteClassB: Step 2"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test hook (empty implementation)
    result = classB.Hook()
    expected = ""
    if result != expected {
        t.Errorf("Expected empty string, got '%s'", result)
    }
    
    // Test template method
    result = classB.TemplateMethod()
    expected = "ConcreteClassB: Step 1 -> ConcreteClassB: Step 2"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    // Test with ConcreteClassA
    clientA := NewClient(NewConcreteClassA())
    result := clientA.ExecuteAlgorithm()
    expected := "ConcreteClassA: Step 1 -> ConcreteClassA: Step 2 -> ConcreteClassA: Hook"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test with ConcreteClassB
    clientB := NewClient(NewConcreteClassB())
    result = clientB.ExecuteAlgorithm()
    expected = "ConcreteClassB: Step 1 -> ConcreteClassB: Step 2"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

// TestTemplateMethodVariations tests different combinations of operations
func TestTemplateMethodVariations(t *testing.T) {
    testCases := []struct {
        name     string
        class    AbstractClass
        expected string
    }{
        {
            name:     "ConcreteClassA with hook",
            class:    NewConcreteClassA(),
            expected: "ConcreteClassA: Step 1 -> ConcreteClassA: Step 2 -> ConcreteClassA: Hook",
        },
        {
            name:     "ConcreteClassB without hook",
            class:    NewConcreteClassB(),
            expected: "ConcreteClassB: Step 1 -> ConcreteClassB: Step 2",
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            client := NewClient(tc.class)
            result := client.ExecuteAlgorithm()
            if result != tc.expected {
                t.Errorf("Expected '%s', got '%s'", tc.expected, result)
            }
        })
    }
} 