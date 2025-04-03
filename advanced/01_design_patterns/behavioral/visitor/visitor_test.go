package visitor

import (
	"testing"
)

func TestConcreteElements(t *testing.T) {
    elementA := NewConcreteElementA("ElementA")
    elementB := NewConcreteElementB("ElementB")
    
    // Test element names
    if name := elementA.GetName(); name != "ElementA" {
        t.Errorf("Expected name 'ElementA', got '%s'", name)
    }
    
    if name := elementB.GetName(); name != "ElementB" {
        t.Errorf("Expected name 'ElementB', got '%s'", name)
    }
}

func TestConcreteVisitor1(t *testing.T) {
    visitor := NewConcreteVisitor1()
    elementA := NewConcreteElementA("ElementA")
    elementB := NewConcreteElementB("ElementB")
    
    // Test visiting ElementA
    result := visitor.VisitConcreteElementA(elementA)
    expected := "Visitor1: Processing ElementA"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test visiting ElementB
    result = visitor.VisitConcreteElementB(elementB)
    expected = "Visitor1: Processing ElementB"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteVisitor2(t *testing.T) {
    visitor := NewConcreteVisitor2()
    elementA := NewConcreteElementA("ElementA")
    elementB := NewConcreteElementB("ElementB")
    
    // Test visiting ElementA
    result := visitor.VisitConcreteElementA(elementA)
    expected := "Visitor2: Processing ElementA"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test visiting ElementB
    result = visitor.VisitConcreteElementB(elementB)
    expected = "Visitor2: Processing ElementB"
    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestObjectStructure(t *testing.T) {
    structure := NewObjectStructure()
    elementA := NewConcreteElementA("ElementA")
    elementB := NewConcreteElementB("ElementB")
    
    // Add elements
    structure.AddElement(elementA)
    structure.AddElement(elementB)
    
    // Test with Visitor1
    visitor1 := NewConcreteVisitor1()
    results := structure.Accept(visitor1)
    
    expected := []string{
        "Visitor1: Processing ElementA",
        "Visitor1: Processing ElementB",
    }
    
    for i, result := range results {
        if result != expected[i] {
            t.Errorf("Expected '%s', got '%s'", expected[i], result)
        }
    }
    
    // Test with Visitor2
    visitor2 := NewConcreteVisitor2()
    results = structure.Accept(visitor2)
    
    expected = []string{
        "Visitor2: Processing ElementA",
        "Visitor2: Processing ElementB",
    }
    
    for i, result := range results {
        if result != expected[i] {
            t.Errorf("Expected '%s', got '%s'", expected[i], result)
        }
    }
}

func TestClient(t *testing.T) {
    client := NewClient()
    elementA := NewConcreteElementA("ElementA")
    elementB := NewConcreteElementB("ElementB")
    
    // Add elements
    client.AddElement(elementA)
    client.AddElement(elementB)
    
    // Test with different visitors
    testCases := []struct {
        name     string
        visitor  Visitor
        expected []string
    }{
        {
            name:    "Visitor1",
            visitor: NewConcreteVisitor1(),
            expected: []string{
                "Visitor1: Processing ElementA",
                "Visitor1: Processing ElementB",
            },
        },
        {
            name:    "Visitor2",
            visitor: NewConcreteVisitor2(),
            expected: []string{
                "Visitor2: Processing ElementA",
                "Visitor2: Processing ElementB",
            },
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            results := client.ExecuteVisitor(tc.visitor)
            for i, result := range results {
                if result != tc.expected[i] {
                    t.Errorf("Expected '%s', got '%s'", tc.expected[i], result)
                }
            }
        })
    }
} 