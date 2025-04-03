package proxy

import (
	"testing"
	"time"
)

func TestRealSubject(t *testing.T) {
    subject := NewRealSubject("test")
    expected := "RealSubject test request"
    if result := subject.Request(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestProxy(t *testing.T) {
    realSubject := NewRealSubject("test")
    proxy := NewProxy(realSubject, time.Second)
    
    // First request should use real subject
    expected := "RealSubject test request"
    if result := proxy.Request(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Second request should use cache
    if result := proxy.Request(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestProtectionProxy(t *testing.T) {
    realSubject := NewRealSubject("test")
    
    // Test with admin permission
    adminProxy := NewProtectionProxy(realSubject, "admin")
    expected := "RealSubject test request"
    if result := adminProxy.Request(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
    
    // Test with non-admin permission
    userProxy := NewProtectionProxy(realSubject, "user")
    expected = "Access denied"
    if result := userProxy.Request(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestVirtualProxy(t *testing.T) {
    proxy := NewVirtualProxy()
    expected := "RealSubject virtual request"
    if result := proxy.Request(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    realSubject := NewRealSubject("test")
    proxy := NewProxy(realSubject, time.Second)
    client := NewClient(proxy)
    
    expected := "RealSubject test request"
    if result := client.UseSubject(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 