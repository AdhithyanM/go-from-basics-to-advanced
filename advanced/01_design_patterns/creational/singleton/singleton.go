package singleton

import (
	"fmt"
	"sync"
)

// Singleton represents a thread-safe singleton instance that stores data.
// It uses a mutex to ensure safe concurrent access to its data.
type Singleton struct {
	// data stores key-value pairs in a map
	data map[string]interface{}
	// mu is a read-write mutex to protect concurrent access to data
	mu sync.RWMutex
}

// instance holds the single instance of Singleton
// once ensures that the instance is created only once
var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the singleton instance.
// It uses sync.Once to ensure thread-safe initialization.
// The first call to GetInstance creates the instance,
// subsequent calls return the existing instance.
func GetInstance() *Singleton {
	// Do ensures that the initialization function is executed only once
	once.Do(func() {
		// Initialize the singleton instance with an empty map
		instance = &Singleton{
			data: make(map[string]interface{}),
		}
	})
	return instance
}

// Set adds or updates a key-value pair in the singleton's data.
// It uses a write lock to ensure thread-safe access.
func (s *Singleton) Set(key string, value interface{}) {
	// Lock the mutex for writing
	s.mu.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer s.mu.Unlock()
	// Store the value in the map
	s.data[key] = value
}

// Get retrieves a value by key from the singleton's data.
// It uses a read lock to allow concurrent reads.
// Returns the value and a boolean indicating if the key exists.
func (s *Singleton) Get(key string) (interface{}, bool) {
	// Lock the mutex for reading
	s.mu.RLock()
	// Ensure the mutex is unlocked when the function returns
	defer s.mu.RUnlock()
	// Retrieve the value from the map
	value, exists := s.data[key]
	return value, exists
}

// Delete removes a key-value pair from the singleton's data.
// It uses a write lock to ensure thread-safe access.
func (s *Singleton) Delete(key string) {
	// Lock the mutex for writing
	s.mu.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer s.mu.Unlock()
	// Remove the key from the map
	delete(s.data, key)
}

// Clear removes all data from the singleton.
// It uses a write lock to ensure thread-safe access.
func (s *Singleton) Clear() {
	// Lock the mutex for writing
	s.mu.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer s.mu.Unlock()
	// Create a new empty map
	s.data = make(map[string]interface{})
}

// String returns a string representation of the singleton's data.
// It uses a read lock to allow concurrent reads.
func (s *Singleton) String() string {
	// Lock the mutex for reading
	s.mu.RLock()
	// Ensure the mutex is unlocked when the function returns
	defer s.mu.RUnlock()
	// Return the string representation of the map
	return fmt.Sprintf("%v", s.data)
} 