package bridge

import "testing"

// TestRefinedAbstraction verifies that the RefinedAbstraction works correctly with different implementors.
func TestRefinedAbstraction(t *testing.T) {
    // Test with ConcreteImplementorA
    implementorA := &ConcreteImplementorA{}
    abstractionA := NewRefinedAbstraction(implementorA)
    
    resultA := abstractionA.Operation()
    expectedA := "RefinedAbstraction: ConcreteImplementorA operation"
    if resultA != expectedA {
        t.Errorf("Expected %s, got %s", expectedA, resultA)
    }
    
    // Test with ConcreteImplementorB
    implementorB := &ConcreteImplementorB{}
    abstractionB := NewRefinedAbstraction(implementorB)
    
    resultB := abstractionB.Operation()
    expectedB := "RefinedAbstraction: ConcreteImplementorB operation"
    if resultB != expectedB {
        t.Errorf("Expected %s, got %s", expectedB, resultB)
    }
}

// TestExtendedAbstraction verifies that the ExtendedAbstraction works correctly with different implementors.
func TestExtendedAbstraction(t *testing.T) {
    // Test with ConcreteImplementorA
    implementorA := &ConcreteImplementorA{}
    abstractionA := NewExtendedAbstraction(implementorA)
    
    resultA := abstractionA.Operation()
    expectedA := "ExtendedAbstraction: ConcreteImplementorA operation"
    if resultA != expectedA {
        t.Errorf("Expected %s, got %s", expectedA, resultA)
    }
    
    // Test with ConcreteImplementorB
    implementorB := &ConcreteImplementorB{}
    abstractionB := NewExtendedAbstraction(implementorB)
    
    resultB := abstractionB.Operation()
    expectedB := "ExtendedAbstraction: ConcreteImplementorB operation"
    if resultB != expectedB {
        t.Errorf("Expected %s, got %s", expectedB, resultB)
    }
}

// TestClient verifies that the Client can use different abstractions correctly.
func TestClient(t *testing.T) {
    client := &Client{}
    
    // Test with RefinedAbstraction and ConcreteImplementorA
    implementorA := &ConcreteImplementorA{}
    abstractionA := NewRefinedAbstraction(implementorA)
    
    resultA := client.UseAbstraction(abstractionA)
    expectedA := "RefinedAbstraction: ConcreteImplementorA operation"
    if resultA != expectedA {
        t.Errorf("Expected %s, got %s", expectedA, resultA)
    }
    
    // Test with ExtendedAbstraction and ConcreteImplementorB
    implementorB := &ConcreteImplementorB{}
    abstractionB := NewExtendedAbstraction(implementorB)
    
    resultB := client.UseAbstraction(abstractionB)
    expectedB := "ExtendedAbstraction: ConcreteImplementorB operation"
    if resultB != expectedB {
        t.Errorf("Expected %s, got %s", expectedB, resultB)
    }
} 