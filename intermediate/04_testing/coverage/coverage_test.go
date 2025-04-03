package coverage

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
    service := NewUserService()
    
    // Test successful creation
    user, err := service.CreateUser("John Doe", 30)
    if err != nil {
        t.Errorf("CreateUser failed: %v", err)
    }
    if user.Name != "John Doe" || user.Age != 30 || !user.IsActive {
        t.Error("User was not created correctly")
    }
    
    // Test negative age
    _, err = service.CreateUser("Jane Doe", -1)
    if err == nil {
        t.Error("Expected error for negative age")
    }
    
    // Test empty name
    _, err = service.CreateUser("", 25)
    if err == nil {
        t.Error("Expected error for empty name")
    }
}

func TestGetUser(t *testing.T) {
    service := NewUserService()
    
    // Create a user first
    user, _ := service.CreateUser("John Doe", 30)
    
    // Test successful retrieval
    retrieved, err := service.GetUser(user.ID)
    if err != nil {
        t.Errorf("GetUser failed: %v", err)
    }
    if retrieved.ID != user.ID {
        t.Error("Retrieved user ID doesn't match")
    }
    
    // Test non-existent user
    _, err = service.GetUser(999)
    if err == nil {
        t.Error("Expected error for non-existent user")
    }
}

func TestUpdateUser(t *testing.T) {
    service := NewUserService()
    
    // Create a user first
    user, _ := service.CreateUser("John Doe", 30)
    
    // Test successful update
    err := service.UpdateUser(user.ID, "John Smith", 31)
    if err != nil {
        t.Errorf("UpdateUser failed: %v", err)
    }
    
    // Verify update
    updated, _ := service.GetUser(user.ID)
    if updated.Name != "John Smith" || updated.Age != 31 {
        t.Error("User was not updated correctly")
    }
    
    // Test invalid updates
    err = service.UpdateUser(user.ID, "", 31)
    if err == nil {
        t.Error("Expected error for empty name")
    }
    
    err = service.UpdateUser(user.ID, "John Smith", -1)
    if err == nil {
        t.Error("Expected error for negative age")
    }
    
    err = service.UpdateUser(999, "John Smith", 31)
    if err == nil {
        t.Error("Expected error for non-existent user")
    }
}

func TestDeleteUser(t *testing.T) {
    service := NewUserService()
    
    // Create a user first
    user, _ := service.CreateUser("John Doe", 30)
    
    // Test successful deletion
    err := service.DeleteUser(user.ID)
    if err != nil {
        t.Errorf("DeleteUser failed: %v", err)
    }
    
    // Verify deletion
    _, err = service.GetUser(user.ID)
    if err == nil {
        t.Error("User was not deleted")
    }
    
    // Test deleting non-existent user
    err = service.DeleteUser(999)
    if err == nil {
        t.Error("Expected error for non-existent user")
    }
}

func TestDeactivateUser(t *testing.T) {
    service := NewUserService()
    
    // Create a user first
    user, _ := service.CreateUser("John Doe", 30)
    
    // Test successful deactivation
    err := service.DeactivateUser(user.ID)
    if err != nil {
        t.Errorf("DeactivateUser failed: %v", err)
    }
    
    // Verify deactivation
    deactivated, _ := service.GetUser(user.ID)
    if deactivated.IsActive {
        t.Error("User was not deactivated")
    }
    
    // Test deactivating non-existent user
    err = service.DeactivateUser(999)
    if err == nil {
        t.Error("Expected error for non-existent user")
    }
}

func TestGetActiveUsers(t *testing.T) {
    service := NewUserService()
    
    // Create some users
    user1, _ := service.CreateUser("John Doe", 30)
    user2, _ := service.CreateUser("Jane Doe", 25)
    service.CreateUser("Bob Smith", 35)
    
    // Deactivate one user
    service.DeactivateUser(user2.ID)
    
    // Get active users
    active := service.GetActiveUsers()
    if len(active) != 2 {
        t.Errorf("Expected 2 active users, got %d", len(active))
    }
    
    // Verify active users
    activeIDs := make(map[int]bool)
    for _, user := range active {
        activeIDs[user.ID] = true
    }
    
    if !activeIDs[user1.ID] {
        t.Error("Expected user1 to be active")
    }
    if activeIDs[user2.ID] {
        t.Error("Expected user2 to be inactive")
    }
} 