package composite

import (
	"testing"
)

func TestLeaf(t *testing.T) {
    leaf := NewLeaf("A")
    if result := leaf.Operation(); result != "Leaf A" {
        t.Errorf("Expected 'Leaf A', got '%s'", result)
    }
}

func TestComposite(t *testing.T) {
    composite := NewComposite("Root")
    leaf1 := NewLeaf("A")
    leaf2 := NewLeaf("B")
    
    composite.Add(leaf1)
    composite.Add(leaf2)
    
    expected := "Composite Root [Leaf A, Leaf B]"
    if result := composite.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    composite.Remove(leaf1)
    expected = "Composite Root [Leaf B]"
    if result := composite.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestNestedComposite(t *testing.T) {
    root := NewComposite("Root")
    branch1 := NewComposite("Branch1")
    branch2 := NewComposite("Branch2")
    
    leaf1 := NewLeaf("A")
    leaf2 := NewLeaf("B")
    leaf3 := NewLeaf("C")
    
    branch1.Add(leaf1)
    branch1.Add(leaf2)
    branch2.Add(leaf3)
    
    root.Add(branch1)
    root.Add(branch2)
    
    expected := "Composite Root [Composite Branch1 [Leaf A, Leaf B], Composite Branch2 [Leaf C]]"
    if result := root.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 