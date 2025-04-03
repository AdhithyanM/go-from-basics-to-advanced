package observer

import (
	"testing"
)

// TestConcreteSubject tests the ConcreteSubject functionality
func TestConcreteSubject(t *testing.T) {
    subject := NewConcreteSubject()
    
    // Test initial state
    if subject.GetState() != "" {
        t.Errorf("Expected empty state, got %s", subject.GetState())
    }
    
    // Test setting state
    subject.SetState("test state")
    if subject.GetState() != "test state" {
        t.Errorf("Expected 'test state', got %s", subject.GetState())
    }
}

// TestConcreteObserverA tests the ConcreteObserverA functionality
func TestConcreteObserverA(t *testing.T) {
    subject := NewConcreteSubject()
    observer := NewConcreteObserverA()
    
    // Test initial state
    if observer.GetState() != "" {
        t.Errorf("Expected empty state, got %s", observer.GetState())
    }
    
    // Test update
    subject.SetState("test state")
    observer.Update(subject)
    if observer.GetState() != "test state" {
        t.Errorf("Expected 'test state', got %s", observer.GetState())
    }
}

// TestConcreteObserverB tests the ConcreteObserverB functionality
func TestConcreteObserverB(t *testing.T) {
    subject := NewConcreteSubject()
    observer := NewConcreteObserverB()
    
    // Test initial state
    if observer.GetState() != "" {
        t.Errorf("Expected empty state, got %s", observer.GetState())
    }
    
    // Test update
    subject.SetState("test state")
    observer.Update(subject)
    if observer.GetState() != "test state" {
        t.Errorf("Expected 'test state', got %s", observer.GetState())
    }
}

// TestObserverRegistration tests observer registration and removal
func TestObserverRegistration(t *testing.T) {
    subject := NewConcreteSubject()
    observerA := NewConcreteObserverA()
    observerB := NewConcreteObserverB()
    
    // Test registration
    subject.RegisterObserver(observerA)
    subject.RegisterObserver(observerB)
    
    // Test notification
    subject.SetState("test state")
    if observerA.GetState() != "test state" {
        t.Errorf("Expected 'test state' for observerA, got %s", observerA.GetState())
    }
    if observerB.GetState() != "test state" {
        t.Errorf("Expected 'test state' for observerB, got %s", observerB.GetState())
    }
    
    // Test removal
    subject.RemoveObserver(observerA)
    subject.SetState("new state")
    if observerA.GetState() != "test state" {
        t.Errorf("Expected 'test state' for removed observerA, got %s", observerA.GetState())
    }
    if observerB.GetState() != "new state" {
        t.Errorf("Expected 'new state' for observerB, got %s", observerB.GetState())
    }
}

// TestClient tests the Client functionality
func TestClient(t *testing.T) {
    client := NewClient()
    observerA := NewConcreteObserverA()
    observerB := NewConcreteObserverB()
    
    // Test adding observers
    client.AddObserver(observerA)
    client.AddObserver(observerB)
    
    // Test setting state
    client.SetState("test state")
    if client.GetObserverState(0) != "test state" {
        t.Errorf("Expected 'test state' for observerA, got %s", client.GetObserverState(0))
    }
    if client.GetObserverState(1) != "test state" {
        t.Errorf("Expected 'test state' for observerB, got %s", client.GetObserverState(1))
    }
    
    // Test removing observer
    client.RemoveObserver(observerA)
    client.SetState("new state")
    if client.GetObserverState(0) != "" {
        t.Errorf("Expected empty state for removed observerA, got %s", client.GetObserverState(0))
    }
    if client.GetObserverState(1) != "new state" {
        t.Errorf("Expected 'new state' for observerB, got %s", client.GetObserverState(1))
    }
} 