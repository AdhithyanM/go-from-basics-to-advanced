package prototype

import "testing"

// TestConcretePrototype1 verifies that ConcretePrototype1 cloning works correctly.
func TestConcretePrototype1(t *testing.T) {
    // Create a prototype
    prototype := NewConcretePrototype1("test")
    prototype.SetData("key1", "value1")
    prototype.SetData("key2", "value2")
    
    // Clone the prototype
    clone := prototype.Clone().(*ConcretePrototype1)
    
    // Verify that the clone is a different object
    if prototype == clone {
        t.Error("Expected different objects")
    }
    
    // Verify that the name was copied
    if clone.name != "test" {
        t.Errorf("Expected name 'test', got %s", clone.name)
    }
    
    // Verify that the data was copied
    if value, exists := clone.data["key1"]; !exists || value != "value1" {
        t.Error("Expected key1 to be copied")
    }
    
    if value, exists := clone.data["key2"]; !exists || value != "value2" {
        t.Error("Expected key2 to be copied")
    }
    
    // Modify the clone and verify it doesn't affect the original
    clone.SetData("key3", "value3")
    if _, exists := prototype.data["key3"]; exists {
        t.Error("Expected original to be unaffected")
    }
}

// TestConcretePrototype2 verifies that ConcretePrototype2 cloning works correctly.
func TestConcretePrototype2(t *testing.T) {
    // Create a prototype
    prototype := NewConcretePrototype2("test")
    prototype.AddData(1)
    prototype.AddData(2)
    prototype.AddData(3)
    
    // Clone the prototype
    clone := prototype.Clone().(*ConcretePrototype2)
    
    // Verify that the clone is a different object
    if prototype == clone {
        t.Error("Expected different objects")
    }
    
    // Verify that the name was copied
    if clone.name != "test" {
        t.Errorf("Expected name 'test', got %s", clone.name)
    }
    
    // Verify that the data was copied
    if len(clone.data) != 3 {
        t.Errorf("Expected 3 elements, got %d", len(clone.data))
    }
    
    for i, value := range clone.data {
        if value != i+1 {
            t.Errorf("Expected %d, got %d", i+1, value)
        }
    }
    
    // Modify the clone and verify it doesn't affect the original
    clone.AddData(4)
    if len(prototype.data) != 3 {
        t.Error("Expected original to be unaffected")
    }
}

// TestPrototypeRegistry verifies that the PrototypeRegistry works correctly.
func TestPrototypeRegistry(t *testing.T) {
    // Create a registry
    registry := NewPrototypeRegistry()
    
    // Add prototypes to the registry
    prototype1 := NewConcretePrototype1("test1")
    prototype1.SetData("key", "value")
    registry.AddPrototype("prototype1", prototype1)
    
    prototype2 := NewConcretePrototype2("test2")
    prototype2.AddData(1)
    registry.AddPrototype("prototype2", prototype2)
    
    // Get and verify prototypes
    clone1 := registry.GetPrototype("prototype1").(*ConcretePrototype1)
    if clone1 == nil {
        t.Error("Expected non-nil clone1")
    }
    
    if clone1.name != "test1" {
        t.Errorf("Expected name 'test1', got %s", clone1.name)
    }
    
    if value, exists := clone1.data["key"]; !exists || value != "value" {
        t.Error("Expected data to be copied")
    }
    
    clone2 := registry.GetPrototype("prototype2").(*ConcretePrototype2)
    if clone2 == nil {
        t.Error("Expected non-nil clone2")
    }
    
    if clone2.name != "test2" {
        t.Errorf("Expected name 'test2', got %s", clone2.name)
    }
    
    if len(clone2.data) != 1 || clone2.data[0] != 1 {
        t.Error("Expected data to be copied")
    }
    
    // Test non-existent prototype
    clone3 := registry.GetPrototype("prototype3")
    if clone3 != nil {
        t.Error("Expected nil for non-existent prototype")
    }
    
    // Test removal
    registry.RemovePrototype("prototype1")
    clone4 := registry.GetPrototype("prototype1")
    if clone4 != nil {
        t.Error("Expected nil after removal")
    }
} 