package adapter

import "testing"

// TestAdapter verifies that the Adapter correctly adapts the Adaptee's interface.
func TestAdapter(t *testing.T) {
    // Create an Adaptee with a specific request
    adaptee := NewAdaptee("specific request")
    
    // Create an Adapter that wraps the Adaptee
    adapter := NewAdapter(adaptee)
    
    // Verify that the Adapter implements the Target interface
    var target Target = adapter
    
    // Test the adapted request
    result := target.Request()
    expected := "Adapter: specific request"
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}

// TestClient verifies that the Client can use the Target interface.
func TestClient(t *testing.T) {
    // Create a Client
    client := &Client{}
    
    // Create an Adaptee and its Adapter
    adaptee := NewAdaptee("client request")
    adapter := NewAdapter(adaptee)
    
    // Use the Target interface through the Client
    result := client.UseTarget(adapter)
    expected := "Adapter: client request"
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}

// TestAdaptee verifies that the Adaptee works correctly on its own.
func TestAdaptee(t *testing.T) {
    // Create an Adaptee
    adaptee := NewAdaptee("test request")
    
    // Test the specific request
    result := adaptee.SpecificRequest()
    expected := "test request"
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}

func TestClassAdapter(t *testing.T) {
    // Create test data
    testData := "Test Data"
    
    // Create class adapter
    classAdapter := NewClassAdapter(testData)
    
    // Create client
    client := &Client{}
    
    // Test adapter
    result := client.UseTarget(classAdapter)
    if result != testData {
        t.Errorf("Expected %s, got %s", testData, result)
    }
}

func TestObjectAdapter(t *testing.T) {
    // Create test data
    testData := "Test Data"
    
    // Create object adapter
    objectAdapter := NewObjectAdapter(testData)
    
    // Create client
    client := &Client{}
    
    // Test adapter
    result := client.UseTarget(objectAdapter)
    if result != testData {
        t.Errorf("Expected %s, got %s", testData, result)
    }
}

func TestAdapterTypes(t *testing.T) {
    // Test that both adapter types implement the Target interface
    var _ Target = (*ClassAdapter)(nil)
    var _ Target = (*ObjectAdapter)(nil)
} 