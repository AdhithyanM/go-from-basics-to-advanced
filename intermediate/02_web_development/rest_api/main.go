package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Item represents a resource in our API
type Item struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}

// In-memory storage
var items = make(map[int]Item)
var nextID = 1

// GET /items - List all items
func handleListItems(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    itemsList := make([]Item, 0, len(items))
    for _, item := range items {
        itemsList = append(itemsList, item)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(itemsList)
}

// GET /items/{id} - Get a single item
func handleGetItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    item, exists := items[id]
    if !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

// POST /items - Create a new item
func handleCreateItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var newItem Item
    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    newItem.ID = nextID
    newItem.CreatedAt = time.Now()
    items[nextID] = newItem
    nextID++

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newItem)
}

// PUT /items/{id} - Update an item
func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if _, exists := items[id]; !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    var updatedItem Item
    if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    updatedItem.ID = id
    updatedItem.CreatedAt = items[id].CreatedAt
    items[id] = updatedItem

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedItem)
}

// DELETE /items/{id} - Delete an item
func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if _, exists := items[id]; !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    delete(items, id)
    w.WriteHeader(http.StatusNoContent)
}

// API versioning middleware
func versionMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        version := r.Header.Get("X-API-Version")
        if version != "v1" {
            http.Error(w, "Unsupported API version", http.StatusBadRequest)
            return
        }
        next(w, r)
    }
}

func main() {
    // Register routes
    http.HandleFunc("/items", versionMiddleware(handleListItems))
    http.HandleFunc("/items/", versionMiddleware(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handleGetItem(w, r)
        case http.MethodPut:
            handleUpdateItem(w, r)
        case http.MethodDelete:
            handleDeleteItem(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))
    http.HandleFunc("/items", versionMiddleware(handleCreateItem))

    // Start server
    fmt.Println("REST API server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
} 