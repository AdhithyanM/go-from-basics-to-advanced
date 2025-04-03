package coverage

import "errors"

// User represents a user in the system
type User struct {
    ID       int
    Name     string
    Age      int
    IsActive bool
}

// UserService handles user-related operations
type UserService struct {
    users map[int]*User
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
    return &UserService{
        users: make(map[int]*User),
    }
}

// CreateUser creates a new user
func (s *UserService) CreateUser(name string, age int) (*User, error) {
    if age < 0 {
        return nil, errors.New("age cannot be negative")
    }
    
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }

    // Generate a new ID (simplified for example)
    id := len(s.users) + 1
    
    user := &User{
        ID:       id,
        Name:     name,
        Age:      age,
        IsActive: true,
    }
    
    s.users[id] = user
    return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*User, error) {
    user, exists := s.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id int, name string, age int) error {
    user, exists := s.users[id]
    if !exists {
        return errors.New("user not found")
    }
    
    if age < 0 {
        return errors.New("age cannot be negative")
    }
    
    if name == "" {
        return errors.New("name cannot be empty")
    }
    
    user.Name = name
    user.Age = age
    return nil
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(id int) error {
    _, exists := s.users[id]
    if !exists {
        return errors.New("user not found")
    }
    
    delete(s.users, id)
    return nil
}

// DeactivateUser sets a user's active status to false
func (s *UserService) DeactivateUser(id int) error {
    user, exists := s.users[id]
    if !exists {
        return errors.New("user not found")
    }
    
    user.IsActive = false
    return nil
}

// GetActiveUsers returns all active users
func (s *UserService) GetActiveUsers() []*User {
    var active []*User
    for _, user := range s.users {
        if user.IsActive {
            active = append(active, user)
        }
    }
    return active
} 