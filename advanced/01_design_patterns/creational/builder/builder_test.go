package builder

import "testing"

// TestBuilder1 verifies that ConcreteBuilder1 builds the product correctly.
func TestBuilder1(t *testing.T) {
    // Create a new builder
    builder := NewConcreteBuilder1()
    
    // Build the product parts
    builder.BuildPartA()
    builder.BuildPartB()
    builder.BuildPartC()
    
    // Get the product
    product := builder.GetProduct()
    
    // Verify the product parts
    if product.GetPartA() != "PartA1" {
        t.Errorf("Expected PartA1, got %s", product.GetPartA())
    }
    
    if product.GetPartB() != "PartB1" {
        t.Errorf("Expected PartB1, got %s", product.GetPartB())
    }
    
    if product.GetPartC() != "PartC1" {
        t.Errorf("Expected PartC1, got %s", product.GetPartC())
    }
}

// TestBuilder2 verifies that ConcreteBuilder2 builds the product correctly.
func TestBuilder2(t *testing.T) {
    // Create a new builder
    builder := NewConcreteBuilder2()
    
    // Build the product parts
    builder.BuildPartA()
    builder.BuildPartB()
    builder.BuildPartC()
    
    // Get the product
    product := builder.GetProduct()
    
    // Verify the product parts
    if product.GetPartA() != "PartA2" {
        t.Errorf("Expected PartA2, got %s", product.GetPartA())
    }
    
    if product.GetPartB() != "PartB2" {
        t.Errorf("Expected PartB2, got %s", product.GetPartB())
    }
    
    if product.GetPartC() != "PartC2" {
        t.Errorf("Expected PartC2, got %s", product.GetPartC())
    }
}

// TestDirector verifies that the Director constructs products correctly.
func TestDirector(t *testing.T) {
    // Create a director with Builder1
    builder1 := NewConcreteBuilder1()
    director := NewDirector(builder1)
    
    // Construct the product
    product1 := director.Construct()
    
    // Verify the product parts
    if product1.GetPartA() != "PartA1" {
        t.Errorf("Expected PartA1, got %s", product1.GetPartA())
    }
    
    if product1.GetPartB() != "PartB1" {
        t.Errorf("Expected PartB1, got %s", product1.GetPartB())
    }
    
    if product1.GetPartC() != "PartC1" {
        t.Errorf("Expected PartC1, got %s", product1.GetPartC())
    }
    
    // Switch to Builder2
    builder2 := NewConcreteBuilder2()
    director.SetBuilder(builder2)
    
    // Construct another product
    product2 := director.Construct()
    
    // Verify the product parts
    if product2.GetPartA() != "PartA2" {
        t.Errorf("Expected PartA2, got %s", product2.GetPartA())
    }
    
    if product2.GetPartB() != "PartB2" {
        t.Errorf("Expected PartB2, got %s", product2.GetPartB())
    }
    
    if product2.GetPartC() != "PartC2" {
        t.Errorf("Expected PartC2, got %s", product2.GetPartC())
    }
} 