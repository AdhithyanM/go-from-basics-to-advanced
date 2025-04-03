package facade

import (
	"testing"
)

func TestSubsystemA(t *testing.T) {
    subsystem := NewSubsystemA("test")
    expected := "SubsystemA test operation"
    if result := subsystem.OperationA(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestSubsystemB(t *testing.T) {
    subsystem := NewSubsystemB("test")
    expected := "SubsystemB test operation"
    if result := subsystem.OperationB(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestSubsystemC(t *testing.T) {
    subsystem := NewSubsystemC("test")
    expected := "SubsystemC test operation"
    if result := subsystem.OperationC(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestFacade(t *testing.T) {
    facade := NewFacade("test")
    expected := `Facade operation:
- SubsystemA test_A operation
- SubsystemB test_B operation
- SubsystemC test_C operation`
    if result := facade.Operation(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestClient(t *testing.T) {
    facade := NewFacade("test")
    client := NewClient(facade)
    expected := `Facade operation:
- SubsystemA test_A operation
- SubsystemB test_B operation
- SubsystemC test_C operation`
    if result := client.UseFacade(); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
} 