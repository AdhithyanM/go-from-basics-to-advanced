package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Basic handler function
func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Home Page!")
}

// Handler with different HTTP methods
func handleMethods(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintf(w, "This is a GET request")
    case http.MethodPost:
        fmt.Fprintf(w, "This is a POST request")
    case http.MethodPut:
        fmt.Fprintf(w, "This is a PUT request")
    case http.MethodDelete:
        fmt.Fprintf(w, "This is a DELETE request")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Handler with query parameters
func handleQuery(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    age := r.URL.Query().Get("age")
    
    if name == "" || age == "" {
        http.Error(w, "Missing parameters", http.StatusBadRequest)
        return
    }
    
    fmt.Fprintf(w, "Hello %s, you are %s years old", name, age)
}

// Handler with JSON response
type Response struct {
    Message string    `json:"message"`
    Time    time.Time `json:"time"`
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "Hello from JSON handler",
        Time:    time.Now(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Handler with form data
func handleForm(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Serve the form
        html := `
        <html>
            <body>
                <form method="POST">
                    <input type="text" name="name" placeholder="Name">
                    <input type="email" name="email" placeholder="Email">
                    <button type="submit">Submit</button>
                </form>
            </body>
        </html>
        `
        fmt.Fprintf(w, html)
        return
    }
    
    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form", http.StatusBadRequest)
            return
        }
        
        name := r.FormValue("name")
        email := r.FormValue("email")
        
        fmt.Fprintf(w, "Received form data: Name=%s, Email=%s", name, email)
    }
}

// Handler with path parameters
func handlePath(w http.ResponseWriter, r *http.Request) {
    // In a real application, you would use a router like gorilla/mux
    // This is a simplified example
    path := r.URL.Path
    fmt.Fprintf(w, "Path: %s", path)
}

// Handler with custom headers
func handleHeaders(w http.ResponseWriter, r *http.Request) {
    // Set custom headers
    w.Header().Set("X-Custom-Header", "Hello")
    w.Header().Set("Content-Type", "text/plain")
    
    // Get request headers
    userAgent := r.Header.Get("User-Agent")
    accept := r.Header.Get("Accept")
    
    fmt.Fprintf(w, "User-Agent: %s\nAccept: %s", userAgent, accept)
}

// Handler with error handling
func handleError(w http.ResponseWriter, r *http.Request) {
    // Simulate an error
    http.Error(w, "This is an error response", http.StatusInternalServerError)
}

func main() {
    // Register handlers
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/methods", handleMethods)
    http.HandleFunc("/query", handleQuery)
    http.HandleFunc("/json", handleJSON)
    http.HandleFunc("/form", handleForm)
    http.HandleFunc("/path/", handlePath)
    http.HandleFunc("/headers", handleHeaders)
    http.HandleFunc("/error", handleError)
    
    // Start server
    fmt.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
} 