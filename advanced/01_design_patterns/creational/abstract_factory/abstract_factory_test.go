package abstract_factory

import "testing"

// TestAbstractFactory verifies that the abstract factory pattern works correctly.
// It tests both factory types and their product creation.
func TestAbstractFactory(t *testing.T) {
    // Test Factory 1
    factory1 := &ConcreteFactory1{}
    
    // Create products using Factory 1
    productA1 := factory1.CreateProductA()
    productB1 := factory1.CreateProductB()
    
    // Verify product creation
    if productA1 == nil || productB1 == nil {
        t.Error("Expected non-nil products")
    }
    
    // Test operations with Factory 1 products
    resultA1 := productA1.OperationA()
    expectedA1 := "ConcreteProductA1 operation"
    if resultA1 != expectedA1 {
        t.Errorf("Expected %s, got %s", expectedA1, resultA1)
    }
    
    resultB1 := productB1.InteractWithA(productA1)
    expectedB1 := "ConcreteProductB1 interacting with ConcreteProductA1 operation"
    if resultB1 != expectedB1 {
        t.Errorf("Expected %s, got %s", expectedB1, resultB1)
    }
    
    // Test Factory 2
    factory2 := &ConcreteFactory2{}
    
    // Create products using Factory 2
    productA2 := factory2.CreateProductA()
    productB2 := factory2.CreateProductB()
    
    // Verify product creation
    if productA2 == nil || productB2 == nil {
        t.Error("Expected non-nil products")
    }
    
    // Test operations with Factory 2 products
    resultA2 := productA2.OperationA()
    expectedA2 := "ConcreteProductA2 operation"
    if resultA2 != expectedA2 {
        t.Errorf("Expected %s, got %s", expectedA2, resultA2)
    }
    
    resultB2 := productB2.InteractWithA(productA2)
    expectedB2 := "ConcreteProductB2 interacting with ConcreteProductA2 operation"
    if resultB2 != expectedB2 {
        t.Errorf("Expected %s, got %s", expectedB2, resultB2)
    }
}

// TestNewFactory verifies the factory creation function.
// It tests creation of different factory types and error handling.
func TestNewFactory(t *testing.T) {
    // Test Factory 1 creation
    factory1 := NewFactory("1")
    if factory1 == nil {
        t.Error("Expected non-nil factory 1")
    }
    
    // Test Factory 1 products
    productA1 := factory1.CreateProductA()
    productB1 := factory1.CreateProductB()
    
    resultA1 := productA1.OperationA()
    expectedA1 := "ConcreteProductA1 operation"
    if resultA1 != expectedA1 {
        t.Errorf("Expected %s, got %s", expectedA1, resultA1)
    }
    
    resultB1 := productB1.InteractWithA(productA1)
    expectedB1 := "ConcreteProductB1 interacting with ConcreteProductA1 operation"
    if resultB1 != expectedB1 {
        t.Errorf("Expected %s, got %s", expectedB1, resultB1)
    }
    
    // Test Factory 2 creation
    factory2 := NewFactory("2")
    if factory2 == nil {
        t.Error("Expected non-nil factory 2")
    }
    
    // Test Factory 2 products
    productA2 := factory2.CreateProductA()
    productB2 := factory2.CreateProductB()
    
    resultA2 := productA2.OperationA()
    expectedA2 := "ConcreteProductA2 operation"
    if resultA2 != expectedA2 {
        t.Errorf("Expected %s, got %s", expectedA2, resultA2)
    }
    
    resultB2 := productB2.InteractWithA(productA2)
    expectedB2 := "ConcreteProductB2 interacting with ConcreteProductA2 operation"
    if resultB2 != expectedB2 {
        t.Errorf("Expected %s, got %s", expectedB2, resultB2)
    }
    
    // Test invalid factory type
    factory3 := NewFactory("3")
    if factory3 != nil {
        t.Error("Expected nil factory for invalid type")
    }
}

// TestClient verifies that the client can use different factories correctly.
func TestClient(t *testing.T) {
    client := &Client{}
    
    // Test with Factory 1
    factory1 := &ConcreteFactory1{}
    resultA1, resultB1 := client.CreateProducts(factory1)
    
    expectedA1 := "ConcreteProductA1 operation"
    if resultA1 != expectedA1 {
        t.Errorf("Expected %s, got %s", expectedA1, resultA1)
    }
    
    expectedB1 := "ConcreteProductB1 interacting with ConcreteProductA1 operation"
    if resultB1 != expectedB1 {
        t.Errorf("Expected %s, got %s", expectedB1, resultB1)
    }
    
    // Test with Factory 2
    factory2 := &ConcreteFactory2{}
    resultA2, resultB2 := client.CreateProducts(factory2)
    
    expectedA2 := "ConcreteProductA2 operation"
    if resultA2 != expectedA2 {
        t.Errorf("Expected %s, got %s", expectedA2, resultA2)
    }
    
    expectedB2 := "ConcreteProductB2 interacting with ConcreteProductA2 operation"
    if resultB2 != expectedB2 {
        t.Errorf("Expected %s, got %s", expectedB2, resultB2)
    }
} 