package flyweight

import (
	"testing"
)

func TestConcreteFlyweight(t *testing.T) {
    flyweight := NewConcreteFlyweight("shared")
    expected := "Intrinsic: shared, Extrinsic: unique"
    if result := flyweight.Operation("unique"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestUnsharedConcreteFlyweight(t *testing.T) {
    flyweight := NewUnsharedConcreteFlyweight("all")
    expected := "Unshared: all"
    if result := flyweight.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestFlyweightFactory(t *testing.T) {
    factory := NewFlyweightFactory()
    
    // First request for a flyweight
    flyweight1 := factory.GetFlyweight("shared")
    expected := "Intrinsic: shared, Extrinsic: unique1"
    if result := flyweight1.Operation("unique1"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Second request for the same flyweight
    flyweight2 := factory.GetFlyweight("shared")
    expected = "Intrinsic: shared, Extrinsic: unique2"
    if result := flyweight2.Operation("unique2"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Verify that both variables point to the same object
    if flyweight1 != flyweight2 {
        t.Error("Expected the same flyweight instance")
    }
}

func TestClient(t *testing.T) {
    factory := NewFlyweightFactory()
    client := NewClient(factory)
    
    // Test shared flyweight
    expected := "Intrinsic: shared, Extrinsic: unique"
    if result := client.UseFlyweight("shared", "unique"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test unshared flyweight
    expected = "Unshared: all"
    if result := client.UseUnsharedFlyweight("all"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 