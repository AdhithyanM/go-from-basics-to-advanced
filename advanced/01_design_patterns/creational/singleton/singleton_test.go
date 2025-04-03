package singleton

import (
	"fmt"
	"testing"
	"time"
)

// TestSingleton verifies that the singleton pattern works correctly.
// It checks that multiple calls to GetInstance return the same instance
// and that the instance's state is shared.
func TestSingleton(t *testing.T) {
	// Get the first instance
	instance1 := GetInstance()
	
	// Get the second instance
	instance2 := GetInstance()
	
	// Verify that both instances are the same
	if instance1 != instance2 {
		t.Error("Expected same instance")
	}
	
	// Test that state is shared between instances
	instance1.Set("key", "value")
	if value, exists := instance2.Get("key"); !exists || value != "value" {
		t.Error("Expected same state")
	}
}

// TestSingletonConcurrency tests the singleton's thread safety
// by performing concurrent operations from multiple goroutines.
func TestSingletonConcurrency(t *testing.T) {
	// Get the singleton instance
	instance := GetInstance()
	
	// Launch 100 goroutines to perform concurrent writes
	for i := 0; i < 100; i++ {
		go func(i int) {
			// Each goroutine sets a unique key-value pair
			instance.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	
	// Wait for all goroutines to complete
	time.Sleep(100 * time.Millisecond)
	
	// Verify that all values were set correctly
	for i := 0; i < 100; i++ {
		if value, exists := instance.Get(fmt.Sprintf("key%d", i)); !exists || value != i {
			t.Errorf("Expected value %d for key%d", i, i)
		}
	}
}

// TestSingletonClear verifies that the Clear method
// removes all data from the singleton.
func TestSingletonClear(t *testing.T) {
	// Get the singleton instance
	instance := GetInstance()
	
	// Add some test data
	instance.Set("key1", "value1")
	instance.Set("key2", "value2")
	
	// Clear all data
	instance.Clear()
	
	// Verify that the data was cleared
	if _, exists := instance.Get("key1"); exists {
		t.Error("Expected key1 to be deleted")
	}
	if _, exists := instance.Get("key2"); exists {
		t.Error("Expected key2 to be deleted")
	}
}

// TestSingletonDelete verifies that the Delete method
// removes a specific key-value pair from the singleton.
func TestSingletonDelete(t *testing.T) {
	// Get the singleton instance
	instance := GetInstance()
	
	// Add test data
	instance.Set("key", "value")
	
	// Delete the key
	instance.Delete("key")
	
	// Verify that the key was deleted
	if _, exists := instance.Get("key"); exists {
		t.Error("Expected key to be deleted")
	}
}

// TestSingletonString verifies that the String method
// returns a proper string representation of the singleton's data.
func TestSingletonString(t *testing.T) {
	// Get the singleton instance
	instance := GetInstance()
	
	// Add test data
	instance.Set("key1", "value1")
	instance.Set("key2", "value2")
	
	// Get the string representation
	str := instance.String()
	
	// Verify the string representation
	expected := "map[key1:value1 key2:value2]"
	if str != expected {
		t.Errorf("Expected %s, got %s", expected, str)
	}
} 