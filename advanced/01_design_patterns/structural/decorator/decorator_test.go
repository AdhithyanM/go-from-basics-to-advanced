package decorator

import (
	"testing"
)

func TestConcreteComponent(t *testing.T) {
    component := NewConcreteComponent("A")
    expected := "ConcreteComponent A"
    if result := component.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteDecoratorA(t *testing.T) {
    component := NewConcreteComponent("A")
    decorator := NewConcreteDecoratorA(component, "state1")
    
    expected := "ConcreteComponent A with ConcreteDecoratorA state1"
    if result := decorator.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestConcreteDecoratorB(t *testing.T) {
    component := NewConcreteComponent("A")
    decorator := NewConcreteDecoratorB(component, func(s string) string {
        return "[" + s + "]"
    })
    
    expected := "[ConcreteComponent A]"
    if result := decorator.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestMultipleDecorators(t *testing.T) {
    component := NewConcreteComponent("A")
    decoratorA := NewConcreteDecoratorA(component, "state1")
    decoratorB := NewConcreteDecoratorB(decoratorA, func(s string) string {
        return "{" + s + "}"
    })
    
    expected := "{ConcreteComponent A with ConcreteDecoratorA state1}"
    if result := decoratorB.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 