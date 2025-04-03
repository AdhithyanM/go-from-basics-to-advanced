package iterator

import "testing"

func TestConcreteIterator(t *testing.T) {
    collection := []interface{}{1, 2, 3}
    iterator := NewConcreteIterator(collection)

    // Test HasNext
    if !iterator.HasNext() {
        t.Error("Expected HasNext to return true")
    }

    // Test Next
    if item := iterator.Next(); item != 1 {
        t.Errorf("Expected 1, got %v", item)
    }

    if item := iterator.Next(); item != 2 {
        t.Errorf("Expected 2, got %v", item)
    }

    if item := iterator.Next(); item != 3 {
        t.Errorf("Expected 3, got %v", item)
    }

    // Test end of collection
    if iterator.HasNext() {
        t.Error("Expected HasNext to return false")
    }

    if item := iterator.Next(); item != nil {
        t.Errorf("Expected nil, got %v", item)
    }
}

func TestConcreteAggregate(t *testing.T) {
    aggregate := NewConcreteAggregate()
    aggregate.AddItem(1)
    aggregate.AddItem(2)
    aggregate.AddItem(3)

    iterator := aggregate.CreateIterator()
    expected := []interface{}{1, 2, 3}
    index := 0

    for iterator.HasNext() {
        if item := iterator.Next(); item != expected[index] {
            t.Errorf("Expected %v, got %v", expected[index], item)
        }
        index++
    }
}

func TestClient(t *testing.T) {
    client := NewClient()
    client.AddItems(1, 2, 3)

    result := client.Iterate()
    expected := []interface{}{1, 2, 3}

    if len(result) != len(expected) {
        t.Errorf("Expected length %d, got %d", len(expected), len(result))
    }

    for i, item := range result {
        if item != expected[i] {
            t.Errorf("Expected %v, got %v", expected[i], item)
        }
    }
} 