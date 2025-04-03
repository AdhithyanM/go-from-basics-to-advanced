package factory_method

import "testing"

// TestFactoryMethod verifies that the factory method pattern works correctly.
// It tests both creator types and their product creation.
func TestFactoryMethod(t *testing.T) {
    // Test Creator A
    creatorA := &ConcreteCreatorA{}
    // Create a product using Creator A
    productA := creatorA.FactoryMethod("ProductA")
    
    // Verify product creation
    if productA == nil {
        t.Error("Expected non-nil product")
    }
    
    // Verify product name
    if productA.GetName() != "ProductA" {
        t.Errorf("Expected name ProductA, got %s", productA.GetName())
    }
    
    // Test operation with Creator A
    resultA := creatorA.SomeOperation("ProductA")
    expectedA := "CreatorA: ConcreteProductA operation"
    if resultA != expectedA {
        t.Errorf("Expected %s, got %s", expectedA, resultA)
    }
    
    // Test Creator B
    creatorB := &ConcreteCreatorB{}
    // Create a product using Creator B
    productB := creatorB.FactoryMethod("ProductB")
    
    // Verify product creation
    if productB == nil {
        t.Error("Expected non-nil product")
    }
    
    // Verify product name
    if productB.GetName() != "ProductB" {
        t.Errorf("Expected name ProductB, got %s", productB.GetName())
    }
    
    // Test operation with Creator B
    resultB := creatorB.SomeOperation("ProductB")
    expectedB := "CreatorB: ConcreteProductB operation"
    if resultB != expectedB {
        t.Errorf("Expected %s, got %s", expectedB, resultB)
    }
}

// TestNewCreator verifies the factory function that creates creators.
// It tests creation of different creator types and error handling.
func TestNewCreator(t *testing.T) {
    // Test Creator A creation
    creatorA := NewCreator("A")
    if creatorA == nil {
        t.Error("Expected non-nil creator A")
    }
    
    // Test Creator A operation
    resultA := creatorA.SomeOperation("ProductA")
    expectedA := "CreatorA: ConcreteProductA operation"
    if resultA != expectedA {
        t.Errorf("Expected %s, got %s", expectedA, resultA)
    }
    
    // Test Creator B creation
    creatorB := NewCreator("B")
    if creatorB == nil {
        t.Error("Expected non-nil creator B")
    }
    
    // Test Creator B operation
    resultB := creatorB.SomeOperation("ProductB")
    expectedB := "CreatorB: ConcreteProductB operation"
    if resultB != expectedB {
        t.Errorf("Expected %s, got %s", expectedB, resultB)
    }
    
    // Test invalid creator type
    creatorC := NewCreator("C")
    if creatorC != nil {
        t.Error("Expected nil creator for invalid type")
    }
}

// TestClient verifies that the client can use different creators correctly.
func TestClient(t *testing.T) {
    client := &Client{}
    
    // Test with Creator A
    creatorA := &ConcreteCreatorA{}
    resultA := client.UseCreator(creatorA, "ProductA")
    expectedA := "CreatorA: ConcreteProductA operation"
    if resultA != expectedA {
        t.Errorf("Expected %s, got %s", expectedA, resultA)
    }
    
    // Test with Creator B
    creatorB := &ConcreteCreatorB{}
    resultB := client.UseCreator(creatorB, "ProductB")
    expectedB := "CreatorB: ConcreteProductB operation"
    if resultB != expectedB {
        t.Errorf("Expected %s, got %s", expectedB, resultB)
    }
} 